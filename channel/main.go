package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func consumer(ch chan int) {
	defer wg.Done()
	//for data := range ch {
	//	fmt.Println(data)
	//	time.Sleep(time.Second)
	//}
	for {
		data, ok := <-ch
		fmt.Println(data, ok)
		time.Sleep(time.Second)
		if !ok {
			break
		}
	}
}

func main() {
	var msg chan int
	msg = make(chan int, 1)
	msg <- 1
	wg.Add(1)
	go consumer(msg)
	msg <- 2
	close(msg)
	wg.Wait()
}
