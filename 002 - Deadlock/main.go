package main

// go env -w GO111MODULE=auto

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	fmt.Println("002 - Deadlock")

	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("the sum is %v. \n", v1.value+v2.value)
	}

	var a, b value

	wg.Add(2)

	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}
