package main

import (
	"regexp"
	"strings"
)

type Arg struct {
	// Target entity name
	Entity string

	// Target slice name
	Slice string

	// Exclude field names
	Excludes []string
}

func newArg(rawArgs []string) Arg {
	pattern := `-(\w+)=(\w+)`

	arg := Arg{}
	for _, rawarg := range rawArgs {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(rawarg)
		if len(matches) != 2 {
			panic("invalid arg: " + rawarg)
		}
		key, val := matches[1], matches[2]
		switch key {
		case "entity":
			arg.Entity = val
		case "slice":
			arg.Slice = val
		case "exclude":
			arg.Excludes = strings.Split(val, ",")
		}
	}
	return arg
}

func (a Arg) validate() bool {
	if a.Entity == "" {
		return false
	}
	if a.Slice == "" {
		return false
	}
	return true
}
