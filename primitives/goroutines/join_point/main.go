package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	sayHello := func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}

	wg.Add(1)
	go sayHello("Hello!!!")
	wg.Wait() // This is the join point

}
