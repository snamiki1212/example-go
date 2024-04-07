package main

import "fmt"

type Number interface {
	int64 | float64
}

func sum[N Number](a N, b N) N {
	return a + b
}

type Person interface {
	Work() string
}

func doPrint(p Person) error {
	_, err := fmt.Println(p.Work())
	return err
}

type Worker1 struct{}

func (w Worker1) Work() string {
	return "worker1"
}

type Worker2 struct{}

func (w Worker2) Work() string {
	return "worker2"
}

func main() {

	{ // method list
		w1 := Worker1{}
		doPrint(w1)

		w2 := Worker1{}
		doPrint(w2)
	}

	{ // constraint type
		fmt.Println("%+w", sum(int64(1), int64(1)))
		fmt.Println("%+w", sum(float64(1.0), float64(1.5)))
	}
}
