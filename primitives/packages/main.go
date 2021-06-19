package main

import (
	"fmt"
	"math"

	m "./math"
)

func main() {
	xs := []float64{1.6, 2.11, 3.57, 4.03}
	avg := m.Average(xs)
	fmt.Println(avg)
	fmt.Println(math.Log10(avg))
}
