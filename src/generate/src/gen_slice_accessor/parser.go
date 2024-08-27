package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"slices"
)

// Parse sorce code to own struct.
func parse(args arguments) (data, error) {
	fset := token.NewFileSet()

	// Convert source code to ast
	node, err := parser.ParseFile(fset, args.input, nil, parser.AllErrors)
	if err != nil {
		return data{}, fmt.Errorf("parse error: %w", err)
	}

	// Parse ast
	astFs, err := doParse(node, args)
	if err != nil {
		return data{}, err
	}

	// Convert ast to own struct
	fs := newFields(astFs).exclude(args.fieldNamesToExclude)

	return data{
		fields:    fs,
		pkgName:   getPackageName(node),
		sliceName: args.slice,
	}, nil
}

// Get package name.
func getPackageName(node *ast.File) string { return node.Name.Name }

// Parse ast.
func doParse(node *ast.File, args arguments) ([]*ast.Field, error) {
	objs := node.Scope.Objects
	obj, ok := objs[args.entity]
	if !ok {
		return nil, fmt.Errorf("entity not found: %s", args.entity)
	}

	decl := obj.Decl
	entity, ok := decl.(*ast.TypeSpec)
	if !ok {
		return nil, fmt.Errorf("invalid entity: %s", args.entity)
	}
	ty := entity.Type

	// https://stackoverflow.com/questions/20234342/get-a-simple-string-representation-of-a-struct-field-s-type
	sty, ok := ty.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("invalid type: %T", ty)
	}
	fs := sty.Fields.List
	return fs, nil
}

type (
	// Data from parsed source code and will be used in code generation.
	data struct {
		fields    fields
		pkgName   string
		sliceName string
	}
	fields []field

	// Struct field from entity in source code.
	field struct {
		Name string // field name
		Type string // field type like string,int64...
	}
)

// Constructor for field.
func newField(raw *ast.Field) field {
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
	return field{
		Name: name,
		Type: typeName,
	}
}

// Constructor for fields.
func newFields(raws []*ast.Field) fields {
	fs := make(fields, 0, len(raws))
	for _, raw := range raws {
		fs = append(fs, newField(raw))
	}
	return fs
}

// Exclude fields by name.
func (fs fields) exclude(targets []string) fields {
	return slices.DeleteFunc(fs, func(f field) bool {
		return slices.Contains(targets, f.Name)
	})
}

// Constructor for method name.
func newMethodName(name string) string { return name + "s" }
