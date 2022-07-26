package main

import (
	"fmt"
	"learngo/helloRpc/clientStub"
)

func main() {
	// 1. 建立连接
	client := clientStub.NewHelloServiceClient("tcp", "localhost:8999")
	var reply string
	err := client.Hello("world", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
