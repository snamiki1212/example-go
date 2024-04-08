package strategy

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftval, rightval int) int {
	return o.Operator.Apply(leftval, rightval)
}
