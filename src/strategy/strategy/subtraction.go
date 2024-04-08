package strategy

type Subtraction struct{}

func (a *Subtraction) Apply(l, r int) int {
	return l - r
}
