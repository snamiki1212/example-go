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
				arguments: arguments{entity: "User", slice: "Users"},
				src: `
package user

type User struct {
	UserID string
	Age    int64
}
`,
			},
			want: data{
				pkgName:   "user",
				sliceName: "Users",
				fields:    fields{{Name: "UserID", Type: "string"}, {Name: "Age", Type: "int64"}},
			},
		},
		{
			name: "ok: callback",
			args: args{
				arguments: arguments{entity: "User", slice: "Users"},
				src: `
package user

type User struct {
	callback0 func()
	callback1 func(x string, x2 bool) (y int64, y2 int32)
	callback2 func(string, bool) (int64, int32)
	callback3 func(u1 User) (u2 *User)
	callback4 func(cb1 func(x1 string) (y1 int)) (cb2 func(x2 string) (y2 int))
}
`,
			},
			want: data{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Name: "callback0", Type: "func() ()"},
					{Name: "callback1", Type: "func(x string, x2 bool) (y int64, y2 int32)"},
					{Name: "callback2", Type: "func(string, bool) (int64, int32)"},
					{Name: "callback3", Type: "func(u1 User) (u2 *User)"},
					{Name: "callback4", Type: "func(cb1 func(x1 string) (y1 int)) (cb2 func(x2 string) (y2 int))"},
				},
			},
		},
		{
			name: "ok: exclude fields",
			args: args{
				arguments: arguments{entity: "User", slice: "Users", fieldNamesToExclude: []string{"Age"}},
				src: `
package user

type User struct {
	UserID string
	Age    int64
}
`,
			},
			want: data{
				pkgName:   "user",
				sliceName: "Users",
				fields:    fields{{Name: "UserID", Type: "string"}},
			},
		},
		{
			name: "ng: invalid src code: syntax error",
			args: args{
				arguments: arguments{entity: "User", slice: "Users"},
				src: `
package user

type User struct {
	UserID string
}
hogehoge // syntax error
`,
			},
			wantErr: true,
		},
		{
			name: "ng: invalid src code: not found package name",
			args: args{
				arguments: arguments{entity: "User", slice: "Users"},
				src: `
// no package name
type User struct {
	UserID string
}
`,
			},
			wantErr: true,
		},
		{
			name: "ng: invalid arguments: not found entity",
			args: args{
				arguments: arguments{entity: "INVALID_ENTITY", slice: "Users"},
				src: `
package user
type User struct {
	UserID string
}
`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name != "ok: callback" {
				t.Skip()
			}
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
