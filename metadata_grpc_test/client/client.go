package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"learngo/grpc_test/proto"
	"time"
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
	md := metadata.Pairs(
		"timestamp", time.Now().Format("2006-01-02 15:04:05"),
		"user", "admin",
		"user", "xiangnan",
		"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
	)
	//md := metadata.New(map[string]string{
	//	"timestamp": time.Now().Format(time.RFC3339),
	//	"user":      "xiangnan",
	//})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "world"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
