package main

func main() {
	// Handle arguments
	args, err := newArgs()
	if err != nil {
		panic(err)
	}

	// Parse source code
	data, err := parse(args, reader)
	if err != nil {
		panic(err)
	}

	// Generate code
	txt, err := generate(data)
	if err != nil {
		panic(err)
	}

	// Write to output file
	err = write(args.output, txt)
	if err != nil {
		panic(err)
	}
}
