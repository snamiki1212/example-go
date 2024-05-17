package service_a

import (
	"fmt"
)

func Run(req *RunReq) {
	req.OnStart()
	fmt.Println("service_a")
}

type RunReq struct {
	OnStart func() bool
}
