package main

import (
	"go/ast"
	"slices"
)

type (
	Fields []Field

	// Struct field from entity in source code
	Field struct {
		Name string
		Type string
	}
)

func newField(raw *ast.Field) Field {
	name := raw.Names[0].Name
	typeName := func() string {
		switch tt := raw.Type.(type) {
		case *ast.Ident:
			return tt.Name
		case *ast.StarExpr:
			return "*" + tt.X.(*ast.Ident).Name
		}
		return "<invalid-type-name>"
	}()
	return Field{
		Name: name,
		Type: typeName,
	}
}

func newFields(raws []*ast.Field) Fields {
	fields := make(Fields, 0, len(raws))
	for _, raw := range raws {
		fields = append(fields, newField(raw))
	}
	return fields
}

func (fs Fields) exclude(targets []string) Fields {
	return slices.DeleteFunc(fs, func(f Field) bool {
		return slices.Contains(targets, f.Name)
	})
}

func NewMethodName(fieldName string) string { return fieldName + "s" }
