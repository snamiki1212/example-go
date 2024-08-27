package main

import (
	"os"
)

func main() {
	// Handle arguments
	args, err := newArgs(os.Args[1:]) // [0] is not args
	if err != nil {
		panic(err)
	}

	// Parse source code
	data := parse(args)

	// Generate code
	txt := generate(data)

	// Write to output file
	if err := write(args.output, txt); err != nil {
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
