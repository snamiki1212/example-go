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
	node, err := parser.ParseFile(fset, args.input, nil, parser.AllErrors)
	if err != nil {
		return data{}, fmt.Errorf("parse error: %w", err)
	}

	// from stg ast to own Fields
	objs := node.Scope.Objects
	obj, ok := objs[args.entity]
	if !ok {
		return data{}, fmt.Errorf("entity not found: %s", args.entity)
	}

	decl := obj.Decl

	entity, ok := decl.(*ast.TypeSpec)
	if !ok {
		return data{}, fmt.Errorf("invalid entity: %s", args.entity)
	}
	ty := entity.Type

	// https://stackoverflow.com/questions/20234342/get-a-simple-string-representation-of-a-struct-field-s-type
	sty, ok := ty.(*ast.StructType)
	if !ok {
		return data{}, fmt.Errorf("invalid type: %T", ty)
	}
	fields := sty.Fields.List

	infos := newInfos(fields).exclude(args.fieldNamesToExclude)

	return data{
		infos:     infos,
		pkgName:   node.Name.Name, // get package name
		sliceName: args.slice,
	}, nil
}
