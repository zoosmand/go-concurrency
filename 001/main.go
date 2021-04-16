package main

// go env -w GO111MODULE=auto

import (
	"fmt"
	"time"
)

func main() {

	var data int

	go func() {
		data++
	}()
	time.Sleep(1 * time.Second) // Very bad idea !!!

	if data == 0 {
		fmt.Printf("the value is %v. \n", data)
	}
}
