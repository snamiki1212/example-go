package main

import (
	"di/service_a"
	"di/service_b"
)

func main() {
	service_a.Run(func() bool {
		service_b.Run()
		return true
	})
}
