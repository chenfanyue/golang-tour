package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = true
	// MaxInt uint64 = 1 << 64 - 1
	MaxInt uint = 1 << 64 - 1
	complex complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("type: %T, value: %v\n", ToBe, ToBe)
	fmt.Printf("type: %T, value: %v\n", MaxInt, MaxInt)
	fmt.Printf("type: %T, value: %v\n", complex, complex)
}
