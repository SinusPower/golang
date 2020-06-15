package main

import (
	"go/ast"
	"go/token"
	"regexp"
	"strings"
)

type (
	data struct {
		PackageName string
		Functions   []function
	}

	function struct {
		VarName  string
		TypeName string
		Fields   []field
	}

	field struct {
		Name  string
		Type  string
		Rules []rule
	}

	rule struct {
		Type   string
		String string
	}
)

func getData(f *ast.File) data {
	outData := data{
		PackageName: f.Name.Name,
		Functions:   []function{},
	}
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
								Name:  fieldName,
								Type:  fieldType,
								Rules: outRules,
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
	return outData
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
	}

	return sb.String()
}

func getRules(tag string) []rule {
	var out []rule

	if tag == "" {
		return out
	}

	re, err := regexp.Compile(`validate:"(.*?)"`)
	if err != nil {
		return out
	}
	match := re.FindStringSubmatch(tag)

	if match == nil {
		return out
	}

	ruleStrs := strings.Split(match[1], "|")
	for _, str := range ruleStrs {
		colonPos := strings.Index(str, ":")
		if colonPos != -1 && colonPos < len(str)-1 {
			rType := str[:colonPos]
			rString := str[colonPos+1:]
			if rType != "" && rString != "" {
				out = append(out, rule{
					Type:   rType,
					String: rString,
				})
			}
		}
	}

	return out
}
