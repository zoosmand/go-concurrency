package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"Hello", "Morning", "Haiduk"} {
		wg.Add(1)

		go func(salut string) {
			defer wg.Done()
			fmt.Println(salut)
		}(salutation)
	}
	wg.Wait()
}
