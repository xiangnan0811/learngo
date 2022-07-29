package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngo/grpc_test/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "world"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
