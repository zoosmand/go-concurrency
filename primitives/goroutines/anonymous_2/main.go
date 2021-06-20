package main

import (
	"fmt"
	"time"
)

func main() {
	sayHello := func(msg string) {
		fmt.Println(msg)
	}

	go sayHello("HELLO!!!")

	// There is no join point, so sleep starts a race condition
	time.Sleep(time.Millisecond * 10)
}
