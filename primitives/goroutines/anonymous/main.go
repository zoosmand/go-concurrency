package main

import (
	"fmt"
	"time"
)

func main() {
	go func(msg string) {
		fmt.Println(msg)
	}("HELLO!!!")

	// There is no join point, so sleep starts a race condition
	time.Sleep(time.Millisecond * 10)
}
