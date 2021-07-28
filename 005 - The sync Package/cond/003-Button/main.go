package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscibe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()

		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	subscibe(button.Clicked, func() {
		fmt.Println("Maximum window.")
		clickRegistered.Done()
	})

	subscibe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscibe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()

	// button.Clicked.Signal()
	// time.Sleep(1 * time.Second)

	// button.Clicked.Signal()
	// time.Sleep(1 * time.Second)

	// button.Clicked.Signal()
	// time.Sleep(1 * time.Second)

	clickRegistered.Wait()

}
