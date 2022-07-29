package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngo/stream_grpc_test/proto"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	// 服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "xiangnan"})
	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}

	// 客户端流模式
	putS, err := c.PutStream(context.Background())
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		_ = putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("%d%v", i, time.Now().Unix()),
		})
		time.Sleep(time.Second)
	}

	// 双向流模式
	allS, err := c.AllStream(context.Background())
	if err != nil {
		panic(err)
	}
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
			fmt.Println("收到服务端的数据：", a.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			err := allS.Send(&proto.StreamReqData{
				Data: fmt.Sprintf("客户端 -> %v", time.Now().Unix()),
			})
			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
