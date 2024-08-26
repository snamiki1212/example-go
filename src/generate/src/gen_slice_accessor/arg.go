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
}

func newArgs(rawArgs []string) args {
	if isDebug {
		return args{
			entity:              "User",
			slice:               "Users",
			fieldNamesToExclude: []string{"Posts"},
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
