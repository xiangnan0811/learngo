package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func consumer(ch chan int) {
	defer wg.Done()
	data := <-ch
	fmt.Println(data)
}

func main() {
	var msg chan int
	msg = make(chan int, 1)
	msg <- 1
	wg.Add(1)
	go consumer(msg)
	msg <- 2
	wg.Wait()
}
