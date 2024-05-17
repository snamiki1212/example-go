package service_a

import (
	"fmt"
)

func Echo() {
	fmt.Println("echo service_a")
}

func Run(req *RunReq) {
	req.OnStart()
	Echo()
}

type RunReq struct {
	OnStart func() bool
}
