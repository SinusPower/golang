package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

const outFileNameSuffix string = "_validation_generated"

type structField struct {
	Name string
	Type string
	Tag  string
}

var (
	ErrCannotReadFile        = errors.New("can not read file")
	ErrCannotParseFile       = errors.New("can not parse file")
	ErrCannotWriteFile       = errors.New("can not write file")
	ErrCannotBuildCode       = errors.New("can not build code")
	ErrCannotBuildFunc       = errors.New("can not build func")
	ErrCannotFormatOutSource = errors.New("can not format source")
)

func Generate(sourcePath string) error {
	sourceBytes, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrCannotReadFile, err)
	}

	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, sourcePath, sourceBytes, 0)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrCannotParseFile, err)
	}

	outSourceCode, err := buildSource(f)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrCannotBuildCode, err)
	}

	var extension string
	sepIndex := strings.LastIndex(sourcePath, string(os.PathSeparator))
	dotIndex := strings.LastIndex(sourcePath, ".")
	sourceFileName := sourcePath
	if dotIndex != -1 && dotIndex > sepIndex { // some folder can have name such "foo.bar"
		sourceFileName = sourcePath[:dotIndex]
		extension = sourcePath[dotIndex:]
	}

	outFileName := sourceFileName + outFileNameSuffix + extension
	err = ioutil.WriteFile(outFileName, []byte(outSourceCode), 0644)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrCannotWriteFile, err)
	}

	return nil
}

// buildSource builds validation source code string from input AST tree
func buildSource(f *ast.File) (string, error) {
	var bb bytes.Buffer
	writeCommon(&bb, f.Name.Name)

	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)   // generic declaration (import, const, type, var)
		if ok && gd.Tok == token.TYPE { // this is a type declaration (maybe with "(...)")
			for _, spec := range gd.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					if st, ok := ts.Type.(*ast.StructType); ok && st.Fields.List != nil { // struct declaration with fields
						structName := ts.Name.Name
						structFields := make([]structField, 0, st.Fields.NumFields())
						for _, field := range st.Fields.List {
							var fieldName string
							if field.Names != nil {
								fieldName = field.Names[0].Name
							}
							var fieldType string
							if ident, ok := field.Type.(*ast.Ident); ok {
								fieldType = ident.Name
							}
							var fieldTag string
							if field.Tag != nil {
								fieldTag = field.Tag.Value
							}

							if fieldName == "" || fieldType == "" || fieldTag == "" {
								continue
							}

							structFields = append(structFields, structField{
								Name: fieldName,
								Type: fieldType,
								Tag:  fieldTag,
							})
						}
						if len(structFields) > 0 {
							err := writeFunc(&bb, structName, structFields)
							if err != nil {
								return "", fmt.Errorf("%s: %w", ErrCannotBuildFunc, err)
							}
						}
					}
				}
			}
		}
	}

	outBytes, err := format.Source(bb.Bytes())
	if err != nil {
		return "", fmt.Errorf("%s: %w", ErrCannotFormatOutSource, err)
	}
	return string(outBytes), nil
}

func writeCommon(bb *bytes.Buffer, packageName string) {
	bb.WriteString("// Code generated by go-validate tool. DO NOT EDIT.\n")
	bb.WriteString("package " + packageName + "\n\n")
	bb.WriteString(`type ValidationError struct {
		Field string
		Error error
}
`)
}

func writeFunc(bb *bytes.Buffer, structName string, structFields []structField) error {
	signature := funcSignature{
		VarName:  getVarName(structName),
		TypeName: structName,
	}

	validationRules := buildValidationRules(structFields)

	tmplData := templateData{
		Signature:       signature,
		ValidationRules: validationRules,
	}

	templates := make(map[string]*template.Template)
	tmpl, err := template.ParseFiles("templates/signature.tmpl", "templates/func.tmpl")
	if err != nil {
		return err
	}
	templates["func"] = tmpl

	var fBytes bytes.Buffer
	err = templates["func"].ExecuteTemplate(&fBytes, "func", tmplData)
	if err != nil {
		return err
	}

	bb.Write(fBytes.Bytes())
	return nil
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
