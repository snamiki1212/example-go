package service_b

import (
	"fmt"
)

func Echo() {
	fmt.Println("echo service_b")
}

func Run(req *RunReq) {
	req.OnStart()
	Echo()
}

type RunReq struct {
	OnStart func() bool
}
