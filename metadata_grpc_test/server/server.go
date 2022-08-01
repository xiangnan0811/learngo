package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	"learngo/grpc_test/proto"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("no metadata")
	}
	for key, value := range md {
		fmt.Printf("%s -> %v\n", key, value)
	}
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to serve: " + err.Error())
	}
}
