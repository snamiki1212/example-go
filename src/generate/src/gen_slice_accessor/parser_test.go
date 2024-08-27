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
		name    string
		args    args
		want    data
		wantErr bool
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
			want: data{
				fields: fields{
					{
						Name: "UserID",
						Type: "string",
					},
				},
				pkgName:   "user",
				sliceName: "Users",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := newReaderFromString(tt.args.src)
			got, err := parse(tt.args.arguments, reader)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// Construct reader from string.
func newReaderFromString(src string) func(path string) (*ast.File, error) {
	return func(path string) (*ast.File, error) {
		fset := token.NewFileSet()
		noFilePath := "" // not import from file path
		return parser.ParseFile(fset, noFilePath, src, parser.AllErrors)
	}
}
