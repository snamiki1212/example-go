package main

import (
	"flag"
	"fmt"
	"strings"
)

type arguments struct {
	// Target entity name
	entity string

	// Target slice name
	slice string

	// Field names to exclude
	fieldNamesToExclude []string

	// Input file name
	input string

	// Output file name
	output string
}

var (
	entity              = flag.String("entity", "", "target entity name")
	slice               = flag.String("slice", "", "target slice name")
	fieldNamesToExclude = flag.String("exclude", "", "field names to exclude") // TODO: rename to ignore
	input               = flag.String("in", "", "input file name")
	output              = flag.String("out", "", "output file name")
)

func newArgs() (arguments, error) {
	flag.Parse()
	args := arguments{
		entity:              *entity,
		slice:               *slice,
		fieldNamesToExclude: strings.Split(*fieldNamesToExclude, ","),
		input:               *input,
		output:              *output,
	}
	if err := args.validate(); err != nil {
		return arguments{}, err
	}
	return args, nil
}

func (a arguments) validate() error {
	var container []error
	if a.entity == "" {
		container = append(container, fmt.Errorf("entity is required"))
	}
	if a.slice == "" {
		container = append(container, fmt.Errorf("slice is required"))
	}
	if len(container) > 0 {
		return fmt.Errorf("invalid arguments: %v", container)
	}
	return nil
}
