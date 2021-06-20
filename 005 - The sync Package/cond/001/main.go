package main

import "sync"

func main() {

	conditionTrue := func() bool {
		return true
	}

	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock()

	for !conditionTrue() {
		c.Wait()
	}

	c.L.Unlock()
}
