package main

// go env -w GO111MODULE=auto

import (
	"fmt"
)

func main() {

	var data int

	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is %v. \n", data)
	}
}
