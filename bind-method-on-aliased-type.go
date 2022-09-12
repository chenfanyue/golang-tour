package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	res := float64(f)
	if res < 0 {
		return -res
	}
	return res
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(math.Sqrt2, f.Abs())
}
