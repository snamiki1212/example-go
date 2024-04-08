package strategy

type Additional struct{}

func (a *Additional) Apply(l, r int) int {
	return l + r
}
