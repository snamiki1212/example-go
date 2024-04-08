package main

import (
	"fmt"
	"strategy/strategy"
)

func main() {
	var op *strategy.Operation
	var r int

	// additional
	op = &strategy.Operation{Operator: &strategy.Additional{}}
	r = op.Operate(2, 3)
	fmt.Println("additional: result is %w", r)

	// multiplication
	op = &strategy.Operation{Operator: &strategy.Multiplication{}}
	r = op.Operate(2, 3)
	fmt.Println("multiplication: result is %w", r)

	// subtraction
	op = &strategy.Operation{Operator: &strategy.Subtraction{}}
	r = op.Operate(2, 3)
	fmt.Println("subtraction: result is %w", r)
}
