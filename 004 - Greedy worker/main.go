package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Print("\n004 - Greedy Worker")
	fmt.Print("\n-------------------\n\n")

	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greeseWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep((3 * time.Nanosecond))
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep((1 * time.Nanosecond))
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep((1 * time.Nanosecond))
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep((1 * time.Nanosecond))
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Polite worker was able to execute %v work loops\n", count)
	}

	wg.Add(2)
	go greeseWorker()
	go politeWorker()

	wg.Wait()
}
