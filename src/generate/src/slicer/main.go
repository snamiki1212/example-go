package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"slices"
)

// const templatePath = "domain/user/gen/getter.tmpl"
const inputPath = "./src/core/domain/user/user.go"
const outputPath = "./src/core/domain/user/user.gen.go"
const inputFile = "user.go"
const outputFile = "user_gen.go"
const entityName = "User"

const isDebug = true // TODO:

func main() {
	fmt.Println("this is slicer")
	excludeFields := []string{"Posts"}

	// args := os.Args

	// sliceName := "Users"

	// pwd
	pwd := getPWD()

	// Output file
	file, err := os.Create(pwd + "/" + outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Ast
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, pwd+"/"+inputFile, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	if isDebug {
		ast.Print(fset, node) // for debug
	}

	{
		objs := node.Scope.Objects
		obj, ok := objs[entityName]
		if !ok {
			panic("not found entity")
		}

		fmt.Println(obj.Name)
		fmt.Println(obj.Kind)
		decl := obj.Decl
		fmt.Println(decl)

		entity, ok := decl.(*ast.TypeSpec)
		if !ok {
			panic("invalid decl")
		}
		fmt.Println(entity)
		ty := entity.Type

		// https://stackoverflow.com/questions/20234342/get-a-simple-string-representation-of-a-struct-field-s-type
		sty, ok := ty.(*ast.StructType)
		if !ok {
			panic("invalid type")
		}
		rawFields := sty.Fields.List
		fmt.Println(rawFields)

		fields := newFields(rawFields)
		fmt.Println(fields)
		fields = fields.exclude(excludeFields)
		fmt.Println(fields)
		// DEBUG
		fmt.Println(":D")

	}

}

func getPWD() string {
	if isDebug {
		return "/Users/snamiki1212/ghq/github.com/snamiki1212/example-go/src/generate/src/core/domain/user"
	}
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(pwd)
	return pwd
}

type (
	Fields []Field
	Field  struct {
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
