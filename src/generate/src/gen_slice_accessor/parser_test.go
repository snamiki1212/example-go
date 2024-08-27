package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parser(t *testing.T) {
	type args struct {
		src       string
		arguments arguments
	}
	tests := []struct {
		name string
		args args
		want fields
	}{
		{
			name: "ok",
			args: args{
				src: `package user

type User struct {
	UserID string
}
`,
				arguments: arguments{
					entity: "User",
					slice:  "Users",
					input:  "user.go",
					output: "users_accessor_gen.go",
				},
			},
			want: fields{
				{
					Name: "UserID",
					Type: "string",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, err := ready(tt.args.src)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			astFs, err := doParse(node, tt.args.arguments)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			got := newFields(astFs).exclude(tt.args.arguments.fieldNamesToExclude)
			assert.Equal(t, tt.want, got)
		})
	}
}

func ready(src string) (f *ast.File, err error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, "", src, parser.AllErrors)
}
