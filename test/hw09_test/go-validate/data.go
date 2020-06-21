package main

import (
	"go/ast"
	"go/token"
	"regexp"
	"strconv"
	"strings"
)

type (
	data struct {
		PackageName string
		Imports     []string
		Functions   []function
	}

	function struct {
		VarName  string
		TypeName string
		Fields   []field
	}

	field struct {
		Name     string
		Type     string
		BaseType string
		Rules    []rule
	}

	rule struct {
		Type   string
		String string
	}
)

func getData(f *ast.File, src []byte) data {
	outData := data{
		PackageName: f.Name.Name,
		Functions:   []function{},
	}
	baseTypes := getBaseTypes(f, src)
	var outFunctions []function
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)   // generic declaration (import, const, type, var)
		if ok && gd.Tok == token.TYPE { // this is a type declaration (maybe with "(...)")
			for _, spec := range gd.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					if st, ok := ts.Type.(*ast.StructType); ok && st.Fields.List != nil { // struct declaration with fields
						outFields := make([]field, 0, st.Fields.NumFields())
						for _, fld := range st.Fields.List {
							var fieldName string
							if fld.Names != nil {
								fieldName = fld.Names[0].Name
							}
							var fieldType string
							if ident, ok := fld.Type.(*ast.Ident); ok {
								fieldType = ident.Name
							}
							var fieldBaseType string
							if arrayType, ok := fld.Type.(*ast.ArrayType); ok {
								fieldType = "array"
								fieldBaseType = string(src[arrayType.Pos()+1 : arrayType.End()-1])
							}
							t, ok := baseTypes[fieldType]
							if ok {
								fieldType = t
							}
							var fieldTag string
							if fld.Tag != nil {
								fieldTag = fld.Tag.Value
							}

							if fieldName == "" || fieldType == "" || fieldTag == "" {
								continue
							}

							outRules := getRules(fieldTag)
							if len(outRules) == 0 {
								continue
							}

							outFields = append(outFields, field{
								Name:     fieldName,
								Type:     fieldType,
								BaseType: fieldBaseType,
								Rules:    outRules,
							})
						}
						if len(outFields) == 0 {
							continue
						}
						structName := ts.Name.Name
						outFunctions = append(outFunctions, function{
							VarName:  getVarName(structName),
							TypeName: structName,
							Fields:   outFields,
						})
					}
				}
			}
		}
	}
	outData.Functions = outFunctions
	setImports(&outData)
	// fmt.Printf("%+v", outData)
	return outData
}

func getBaseTypes(f *ast.File, src []byte) map[string]string {
	out := make(map[string]string)
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)   // generic declaration (import, const, type, var)
		if ok && gd.Tok == token.TYPE { // this is a type declaration (maybe with "(...)")
			for _, spec := range gd.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					if ident, ok := ts.Type.(*ast.Ident); ok {
						out[ts.Name.Name] = ident.Name
					}
				}
			}
		}
	}
	return out
}

func getVarName(structName string) string {
	if structName == "" {
		return ""
	}

	var sb strings.Builder
	letters := strings.Split(structName, "")
	for _, letter := range letters {
		if strings.ToLower(letter) != letter {
			sb.WriteString(strings.ToLower(letter))
		}
		if sb.Len() == 3 {
			break
		}
	}

	if sb.Len() == 0 {
		sb.WriteString(letters[0])
	}

	return sb.String()
}

func getRules(tag string) []rule {
	var out []rule

	if tag == "" {
		return out
	}

	tag = strings.TrimSpace(tag)

	re, err := regexp.Compile(`validate:(\s*?)"(((\s*?)(.*?)(\s*?))*?)"`)
	if err != nil {
		return out
	}
	match := strings.TrimSpace(re.FindString(tag))

	if match == "" {
		return out
	}

	// sequential trimming to avoid unexpected deletion of necessary characters
	match = strings.Trim(strings.Trim(match, "validte:\r\n\t"), `"`)
	ruleStrs := strings.Split(match, "|")
	for _, str := range ruleStrs {
		str = strings.TrimSpace(str)
		colonPos := strings.Index(str, ":")
		if colonPos != -1 && colonPos < len(str)-1 {
			rType := strings.TrimSpace(str[:colonPos])
			rString := strings.TrimSpace(str[colonPos+1:])
			if rType != "" && rString != "" {
				if rType == "min" || rType == "max" || rType == "len" {
					_, err := strconv.Atoi(rString)
					if err != nil {
						continue
					}
				}
				out = append(out, rule{
					Type:   rType,
					String: rString,
				})
			}
		}
	}

	return out
}

func setImports(d *data) {
	imports := []string{"errors"}
	for _, function := range d.Functions {
		for _, field := range function.Fields {
			if field.Type == "string" || field.Type == "[]string" {
				for _, rule := range field.Rules {
					switch rule.Type {
					case "regexp":
						if !in(imports, "regexp") {
							imports = append(imports, "regexp")
						}
					case "in":
						if !in(imports, "strings") {
							imports = append(imports, "strings")
						}
					}
				}
			}
			if field.Type == "int" || field.Type == "[]int" {
				for _, rule := range field.Rules {
					switch rule.Type {
					case "in":
						if !in(imports, "strings") {
							imports = append(imports, "strings")
						}
						if !in(imports, "strconv") {
							imports = append(imports, "strconv")
						}
					}
				}
			}
		}
	}
	d.Imports = imports
}

func in(array []string, str string) bool {
	var found bool
	for _, s := range array {
		found = str == s
		if found {
			break
		}
	}
	return found
}
