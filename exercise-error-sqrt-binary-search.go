package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v",
		float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	// binary search
	left, right := .0, x
	var z float64
	for left <= right {
		z = (left + right) / 2
		fmt.Println(z)
		if z*z > x {
			right = z
		} else if x-z*z > 1e-10 {
			left = z
		} else {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
