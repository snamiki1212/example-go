package main

import (
	"os"
)

func main() {
	arguments := newArgs(os.Args[1:]) // [0] is not args
	if !arguments.validate() {
		panic("invalid args")
	}
	doMain(arguments)
}

func doMain(arguments args) {
	// Parse source code
	data := parse(arguments)

	// Generate code
	txt := generate(data)

	// Write to output file
	if err := write(arguments.output, txt); err != nil {
		panic(err)
	}
}

// Write to output file
func write(path, txt string) error {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(txt)
	return err
}
