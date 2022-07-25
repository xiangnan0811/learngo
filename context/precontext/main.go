package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var stop = make(chan bool)

func cpuInfo() {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("cpu info end")
			return
		default:
			fmt.Println("cpu info")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	wg.Add(1)
	go cpuInfo()
	time.Sleep(time.Second * 3)
	stop <- true
	wg.Wait()
	fmt.Println("end")
}
