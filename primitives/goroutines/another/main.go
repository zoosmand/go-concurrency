package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()

	// There is no join point, so sleep starts a race condition
	time.Sleep(time.Millisecond * 10)
}

func sayHello() {
	fmt.Println("Hello!!!")
}
