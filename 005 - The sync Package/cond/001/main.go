package main

import (
	"sync"
	"time"
)

func main() {

	conditionTrue := func() bool {
		time.Sleep(time.Millisecond * 5)
		return true
	}

	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock()

	for !conditionTrue() {
		c.Wait()
	}

	c.L.Unlock()
}
