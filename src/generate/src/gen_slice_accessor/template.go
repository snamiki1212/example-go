package main

const templateBody = `
// {{ .Method }}
func (xs {{ .Slices }}) {{ .Method }}() []{{ .Type }} {
	sli := make([]{{ .Type }}, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].{{ .Field }})
	}
	return sli
}
`

// Replace variable from key to value in template.
type TemplateMapper struct {
	Slices string
	Method string
	Type   string
	Field  string
}
