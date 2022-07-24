package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 2)
	timeout2 := make(chan bool, 2)
	go func() {
		time.Sleep(time.Second * 1)
		timeout <- true
	}()
	go func() {
		time.Sleep(time.Second * 3)
		timeout2 <- true
	}()
	select {
	case <-timeout:
		fmt.Println("timeout")
	case <-timeout2:
		fmt.Println("timeout2")
		//default:
		//	fmt.Println("default")
	}
	fmt.Println("end")
}
