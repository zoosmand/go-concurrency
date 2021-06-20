package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "Hello!!!"

	wg.Add(1)
	// gorouting runs at the same address space, so it is possible to change "salitation"
	go func() {
		defer wg.Done()
		salutation = "Welcome"
	}()

	wg.Wait()
	fmt.Println(salutation)
}
