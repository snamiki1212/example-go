package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func parse(args arguments) data {
	// Ast
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, args.input, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// from stg ast to own Fields
	var fields infos
	{
		objs := node.Scope.Objects
		obj, ok := objs[args.entity]
		if !ok {
			panic("not found entity")
		}

		decl := obj.Decl

		entity, ok := decl.(*ast.TypeSpec)
		if !ok {
			panic("invalid decl")
		}
		ty := entity.Type

		// https://stackoverflow.com/questions/20234342/get-a-simple-string-representation-of-a-struct-field-s-type
		sty, ok := ty.(*ast.StructType)
		if !ok {
			panic("invalid type")
		}
		rawFields := sty.Fields.List

		fields = newInfos(rawFields)
		fields = fields.exclude(args.fieldNamesToExclude)
	}
	return data{
		infos:     fields,
		pkgName:   node.Name.Name, // get package name
		sliceName: args.slice,
	}
}
