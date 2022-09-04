package main

import (
	"fmt"
	"math"
)

/*
func Sqrt(x float64) float64 {
	z := 1.0
	for z*z-x > 1e-10 || x-z*z > 1e-10 {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}
*/
func Sqrt(x float64) float64 {
	z := 1.0
	for z*z > x || x-z*z > 1e-10 {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
