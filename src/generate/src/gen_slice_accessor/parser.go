package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"slices"
	"strings"
)

// Parse sorce code to own struct.
func parse(args arguments, reader func(path string) (*ast.File, error)) (data, error) {
	// Convert source code to ast
	node, err := reader(args.input)
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

// Read source code from file.
func reader(path string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, path, nil, parser.AllErrors)
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
	name := safeGetName(raw)
	typeName := func() string {
		switch tt := raw.Type.(type) {
		case *ast.Ident:
			return parseIdent(tt)
		case *ast.StarExpr:
			return parseStarExpr(tt)
		case *ast.FuncType:
			return parseFuncType(tt)
		}
		log.Println("parse error: unknown type")
		return "any" // parse error
	}()
	return field{
		Name: name,
		Type: typeName,
	}
}

// Parse identifier.
func parseIdent(x *ast.Ident) string {
	return x.Name
}

// Parse star expression.
func parseStarExpr(x *ast.StarExpr) string {
	switch tt := x.X.(type) {
	case *ast.Ident:
		return "*" + tt.Name
	default:
		log.Println("parseStarExpr: parse error: unknown type")
		return "any"
	}
}

// Parse function type.
func parseFuncType(x *ast.FuncType) string {
	params := func() string {
		if x != nil && x.Params != nil && x.Params.List != nil {
			return newFields(x.Params.List).display()
		}
		return ""
	}()
	results := func() string {
		if x != nil && x.Results != nil && x.Results.List != nil {
			return newFields(x.Results.List).display()
		}
		return ""
	}()
	return fmt.Sprintf("func(%s) (%s)", params, results)
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

// Display fields.
func (fs fields) display() string {
	if len(fs) == 0 {
		return ""
	}
	var pairs []string
	for _, f := range fs {
		pairs = append(pairs, f.display())
	}
	return strings.Join(pairs, ", ")
}

// Display field.
func (f field) display() string {
	if f.Name == "" {
		return f.Type
	}
	return fmt.Sprintf("%s %s", f.Name, f.Type)
}

// Constructor for method name.
// TODO: use pluralize package: https://github.com/gertd/go-pluralize
func newMethodName(name string) string { return name + "s" }

// Get field name safely.
func safeGetName(raw *ast.Field) string {
	if len(raw.Names) == 0 {
		return ""
	}
	return raw.Names[0].Name
}
