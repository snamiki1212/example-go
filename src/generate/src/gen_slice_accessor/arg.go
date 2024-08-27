package main

import (
	"fmt"
	"regexp"
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

// TODO: Use flag package instead of own logic
func newArgs(rawArgs []string) (arguments, error) {
	if true { // TODO:
		return arguments{
			entity:              "User",
			slice:               "Users",
			fieldNamesToExclude: []string{"Posts"},
			input:               "user.go",
			output:              "user_gen.go",
		}, nil
	}
	pattern := `-(\w+)=(\w+)`
	args := arguments{}
	for _, rawarg := range rawArgs {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(rawarg)
		if len(matches) != 2 {
			return arguments{}, fmt.Errorf("invalid arg: %s", rawarg)
		}
		key, val := matches[1], matches[2]
		switch key {
		case "entity":
			args.entity = val
		case "slice":
			args.slice = val
		case "exclude":
			args.fieldNamesToExclude = strings.Split(val, ",")
		}
	}
	if err := args.validate(); err != nil {
		return arguments{}, err
	}
	return args, nil
}

func (a arguments) validate() error {
	if a.entity == "" {
		return fmt.Errorf("entity is required")
	}
	if a.slice == "" {
		return fmt.Errorf("slice is required")
	}
	return nil
}
