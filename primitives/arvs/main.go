package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	maxp := flag.Int("max", 6, "the max value")
	minp := flag.Int("min", 2, "the min value")

	flag.Parse()

	fmt.Println(rand.Intn(*maxp))
	fmt.Println(*minp)
	fmt.Println(*maxp)
}
