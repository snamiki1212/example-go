package main

import (
	"di/service_a"
	"di/service_b"
	"fmt"
)

func main() {
	//
	fmt.Println("---")
	fmt.Println("running service a...")
	service_a.Run(&service_a.RunReq{
		OnStart: func() bool {
			service_b.Echo()
			return true
		},
	})

	//
	fmt.Println("---")
	fmt.Println("running service b...")
	service_b.Run(&service_b.RunReq{
		OnStart: func() bool {
			service_a.Echo()
			return true
		},
	})
}
