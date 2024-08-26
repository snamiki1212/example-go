package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	arguments := newArgs(os.Args[1:]) // [0] is not args
	if !arguments.validate() {
		panic("invalid args")
	}
	doMain(arguments)
}

// 1. Handle request
// 2. Get ast and new field from src
// 3. Create output / resolve template

func doMain(arguments args) {

	// Ast
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, arguments.input, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// get package name
	pkgName := node.Name.Name

	// from stg ast to own Fields
	var fields Fields
	{
		objs := node.Scope.Objects
		obj, ok := objs[arguments.entity]
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

		fields = newFields(rawFields)
		fields = fields.exclude(arguments.fieldNamesToExclude)
	}

	// Generate code
	txt := generate(pkgName, arguments.slice, fields)

	// Write to output file
	write(arguments.output, txt)
}

// Write to output file
func write(path, txt string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(txt)
	if err != nil {
		panic(err)
	}
}
