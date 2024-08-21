package main

//go:generate go run ./gen.go
import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	// inputPath := "../user.go"
	templatePath := "./getter.tmpl"
	outputPath := "../user.gen.go"
	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	t := template.Must(template.ParseFiles(templatePath))
	replaceRule := map[string]string{
		"REPLACE_ME": ":D",
	} //
	if err := t.Execute(file, replaceRule); err != nil {
		panic(err)
	}
	fmt.Println("generate!!!")
}
