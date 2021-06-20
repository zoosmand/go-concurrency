package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goruting is sleeping...")
		time.Sleep(time.Millisecond * 1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2st goruting is sleeping...")
		time.Sleep(time.Millisecond * 2)
	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")
}
