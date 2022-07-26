package main

import (
	"learngo/helloRpc/serverStub"
	"net"
	"net/rpc"

	"learngo/helloRpc/handler"
)

func main() {
	// 1. 实例化一个rpc服务
	listener, _ := net.Listen("tcp", ":8999")
	// 2. 注册处理逻辑 handler
	err := serverStub.RegisterHelloService(&handler.HelloService{})
	if err != nil {
		panic(err)
	}
	// 3. 启动rpc服务
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
