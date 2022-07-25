package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
	go cpuInfo(ctx)
	go memoryInfo(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
	fmt.Println("end")
}

func cpuInfo(ctx context.Context) {
	defer wg.Done()
	//go memoryInfo(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cpu info end")
			return
		default:
			fmt.Println("cpu info")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func memoryInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("memory info end")
			return
		default:
			fmt.Println("memory info")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
