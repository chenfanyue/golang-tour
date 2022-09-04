package main

import (
	"fmt"
	"math"
)

func pow(x, n, pivot float64) float64 {
	if v := math.Pow(x, n); v < pivot {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, pivot)
	}
	return pivot
}

func main() {
	fmt.Println(
		pow(3, 2, 20),
		pow(3, 3, 20),
	)
}
