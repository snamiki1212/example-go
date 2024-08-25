package gen

//go:generate go run ./gen.go
import (
	"fmt"
	"html/template"
	"os"
)

const templatePath = "domain/user/gen/getter.tmpl"
const outputPath = "domain/user/user.gen.go"

func main() {
	// inputPath := "../user.go"
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
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

func Main() {
	main()
}
