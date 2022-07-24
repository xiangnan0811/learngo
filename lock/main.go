package main

import (
	"fmt"
	"sync"
)

var total int
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		total++
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		total--
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
