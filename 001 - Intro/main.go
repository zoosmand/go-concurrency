package main

// go env -w GO111MODULE=auto

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("001 - Intro")

	var memoryAccess sync.Mutex
	var value int

	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	time.Sleep(1 * time.Second)
	memoryAccess.Lock()

	if value == 0 {
		fmt.Println("the value is 0.")
	} else {
		fmt.Printf("the value is %v. \n", value)
	}

	memoryAccess.Unlock()
}
