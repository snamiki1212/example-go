package main

import (
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
func newArgs(rawArgs []string) arguments {
	if true { // TODO:
		return arguments{
			entity:              "User",
			slice:               "Users",
			fieldNamesToExclude: []string{"Posts"},
			input:               "user.go",
			output:              "user_gen.go",
		}
	}
	pattern := `-(\w+)=(\w+)`
	args := arguments{}
	for _, rawarg := range rawArgs {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(rawarg)
		if len(matches) != 2 {
			panic("invalid arg: " + rawarg)
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
	return args
}

func (a arguments) validate() bool {
	if a.entity == "" {
		return false
	}
	if a.slice == "" {
		return false
	}
	return true
}
