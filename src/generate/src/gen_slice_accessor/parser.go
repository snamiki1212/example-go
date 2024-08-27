package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// Parse sorce code.
// [flow] srcCode -> Ast -> data
func parse(args arguments) (data, error) {
	fset := token.NewFileSet()

	// Convert source code to ast
	node, err := parser.ParseFile(fset, args.input, nil, parser.AllErrors)
	if err != nil {
		return data{}, fmt.Errorf("parse error: %w", err)
	}

	// Parse ast
	fields, err := doParse(node, args)
	if err != nil {
		return data{}, err
	}

	// Convert ast to own struct
	infos := newInfos(fields).exclude(args.fieldNamesToExclude)

	return data{
		infos:     infos,
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
	fields := sty.Fields.List
	return fields, nil
}
