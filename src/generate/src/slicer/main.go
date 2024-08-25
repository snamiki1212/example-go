package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// const templatePath = "domain/user/gen/getter.tmpl"
const inputPath = "./src/core/domain/user/user.go"
const outputPath = "./src/core/domain/user/user.gen.go"
const inputFile = "user.go"
const outputFile = "user_gen.go"

// func main() {
// 	// inputPath := "../user.go"
// 	path, err := os.Getwd()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(path)

// 	// output file
// 	file, err := os.Create(outputPath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	fset := token.NewFileSet()
// 	node, err := parser.ParseFile(fset, inputPath, nil, parser.AllErrors)
// 	if err != nil {
// 		panic(err)
// 	}
// 	ast.Print(fset, node)

// 	// // template
// 	// t := template.Must(template.ParseFiles(templatePath))
// 	// replaceRule := map[string]string{
// 	// 	"REPLACE_ME": ":D",
// 	// } //
// 	// if err := t.Execute(file, replaceRule); err != nil {
// 	// 	panic(err)
// 	// }
// 	fmt.Println("generate!!!")
// }

// func Main() {
// 	main()
// }

func main() {
	fmt.Println("this is slicer")

	// args := os.Args
	// entityName := "User"
	// sliceName := "Users"

	// pwd
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(pwd)

	// output file
	file, err := os.Create(pwd + "/" + outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, pwd+"/"+inputFile, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, node)
}
