package main

// go env -w GO111MODULE=auto

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	fmt.Println("003 - Livelock")

	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprintf(out, "Success!")
			return true
		}

		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32

	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("left", &left, out)
	}

	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("right", &right, out)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() {
			fmt.Println(out.String())
		}()
		defer walking.Done()

		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperatino!", name)
	}

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barnara")
	peopleInHallway.Wait()

	// var wg sync.WaitGroup

	// printSum := func(v1, v2 *value) {
	// 	defer wg.Done()
	// 	v1.mu.Lock()
	// 	defer v1.mu.Unlock()
	// 	time.Sleep(2 * time.Second)
	// 	v2.mu.Lock()
	// 	defer v2.mu.Unlock()

	// 	fmt.Printf("the sum is %v. \n", v1.value+v2.value)
	// }

	// var a, b value

	// wg.Add(2)

	// go printSum(&a, &b)
	// go printSum(&b, &a)

	// wg.Wait()
}
