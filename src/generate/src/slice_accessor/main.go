package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"text/template"
)

const templatePath = "./accessor.tmpl"
const debugPWD = "/Users/snamiki1212/ghq/github.com/snamiki1212/example-go/src/generate/src/core/domain/user"

const isDebug = true // TODO:

func main() {
	pwd := getPWD()
	const inputFile = "user.go"
	const outputFile = "user_gen.go"
	in := pwd + "/" + inputFile
	out := pwd + "/" + outputFile
	doMain(in, out, os.Args[1:]) // [0] is not args
}

// 1. Handle request
// 2. Get ast and new field from src
// 3. Create output / resolve template

func doMain(in, out string, argParams []string) {
	arguments := newArgs(argParams)
	if !arguments.validate() {
		panic("invalid args")
	}

	// Ast
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, in, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	// if isDebug {
	// 	ast.Print(fset, node) // for debug
	// }

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

	// Output file
	output, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer output.Close()
	// append
	_, err = output.WriteString("// Code generated by go generate DO NOT EDIT.\n\n")
	if err != nil {
		panic(err)
	}
	_, err = output.WriteString("package " + pkgName + "\n\n")
	if err != nil {
		panic(err)
	}

	// append templates
	{
		tp, err := template.ParseFiles(templatePath)
		if err != nil {
			panic(err)
		}
		for _, f := range fields {
			data := map[string]string{
				"Slices": arguments.slice,
				"Method": NewMethodName(f.Name),
				"Type":   f.Type,
				"Field":  f.Name,
			}
			err = tp.Execute(output, data)
			if err != nil {
				panic(err)
			}
		}
	}
}

func getPWD() string {
	if isDebug {
		return debugPWD
	}
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return pwd
}
