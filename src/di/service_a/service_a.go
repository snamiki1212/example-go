package service_a

import (
	"fmt"
)

func Run(onStart func() bool) {
	onStart()

	fmt.Println("service_a")
}
