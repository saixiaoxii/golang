package main

import (
	"fmt"
	"sync"
	"time"
)

type Teacher struct {
	Name   string
	Age    int
	School string
}

func test(n int) {
	defer wg.Done()
	for i := n; i < 10; i++ {
		fmt.Println("wuhu", n, i)
		time.Sleep(time.Second)
	}

}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go test(3)
	go test(2)
	wg.Wait()
}
