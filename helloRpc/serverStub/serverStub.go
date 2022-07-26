package serverStub

import (
	"net/rpc"

	"learngo/helloRpc/handler"
)

type HelloService interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(server HelloService) error {
	return rpc.RegisterName(handler.HelloServiceName, server)
}
