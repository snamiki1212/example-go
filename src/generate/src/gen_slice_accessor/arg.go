package main

import (
	"regexp"
	"strings"
)

type args struct {
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
func newArgs(rawArgs []string) args {
	if true { // TODO:
		return args{
			entity:              "User",
			slice:               "Users",
			fieldNamesToExclude: []string{"Posts"},
			input:               "user.go",
			output:              "user_gen.go",
		}
	}
	pattern := `-(\w+)=(\w+)`
	arg := args{}
	for _, rawarg := range rawArgs {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(rawarg)
		if len(matches) != 2 {
			panic("invalid arg: " + rawarg)
		}
		key, val := matches[1], matches[2]
		switch key {
		case "entity":
			arg.entity = val
		case "slice":
			arg.slice = val
		case "exclude":
			arg.fieldNamesToExclude = strings.Split(val, ",")
		}
	}
	return arg
}

func (a args) validate() bool {
	if a.entity == "" {
		return false
	}
	if a.slice == "" {
		return false
	}
	return true
}
