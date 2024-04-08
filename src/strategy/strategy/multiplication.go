package strategy

type Multiplication struct{}

func (a *Multiplication) Apply(l, r int) int {
	return l * r
}
