package main

import (
	"go/ast"
	"slices"
)

type (
	data struct {
		infos     infos
		pkgName   string
		sliceName string
	}
	infos []info

	// Struct field from entity in source code
	info struct {
		Name string
		Type string
	}
)

func newInfo(raw *ast.Field) info {
	name := raw.Names[0].Name
	typeName := func() string {
		switch tt := raw.Type.(type) {
		case *ast.Ident:
			return tt.Name
		case *ast.StarExpr:
			return "*" + tt.X.(*ast.Ident).Name
		}
		return "<invalid-type-name>"
	}()
	return info{
		Name: name,
		Type: typeName,
	}
}

func newInfos(raws []*ast.Field) infos {
	fields := make(infos, 0, len(raws))
	for _, raw := range raws {
		fields = append(fields, newInfo(raw))
	}
	return fields
}

func (fs infos) exclude(targets []string) infos {
	return slices.DeleteFunc(fs, func(f info) bool {
		return slices.Contains(targets, f.Name)
	})
}

func newMethodName(name string) string { return name + "s" }
