package clientStub

import (
	"net/rpc"

	"learngo/helloRpc/handler"
)

type HelloServiceStub struct {
	*rpc.Client
}

func (p *HelloServiceStub) Hello(request string, reply *string) error {
	return p.Call(handler.HelloServiceName+".Hello", request, reply)
}

func NewHelloServiceClient(protocol, address string) HelloServiceStub {
	client, err := rpc.Dial(protocol, address)
	if err != nil {
		panic(err)
	}
	return HelloServiceStub{client}
}
