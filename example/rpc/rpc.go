package main

import (
	"fmt"

	"github.com/Meduzz/disconnected"
	"github.com/Meduzz/disconnected/pkg/event"
	"github.com/Meduzz/rpc"
	"github.com/Meduzz/rpc/encoding"
)

type (
	Request struct {
		Name string `json:"name"`
	}

	Greeting struct {
		Message string `json:"message"`
	}

	ErrorDTO struct {
		Error string `json:"error"`
	}
)

func main() {
	disconnected.RpcServer(encoding.Json(), func(s *event.Server) error {
		return s.HandleRPC("test.greeting", "", Greeter)
	})
}

func Greeter(ctx *rpc.RpcContext) {
	r := &Request{}
	err := ctx.Bind(r)

	if err != nil {
		ctx.Reply(&ErrorDTO{err.Error()})
		return
	}

	res := &Greeting{
		Message: fmt.Sprintf("Hello %s!", r.Name),
	}

	ctx.Reply(res)
}
