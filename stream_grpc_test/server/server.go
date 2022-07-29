package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"

	"learngo/stream_grpc_test/proto"
)

const PORT = ":50052"

type server struct{}

func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		err := res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *server) PutStream(gps proto.Greeter_PutStreamServer) error {
	for {
		a, err := gps.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}
	return nil
}

func (s *server) AllStream(allS proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			a, err := allS.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("收到客户端的数据：", a.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			err := allS.Send(&proto.StreamResData{
				Data: fmt.Sprintf("服务端 -> %v", time.Now().Unix()),
			})
			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
