package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needBig(x float64) float64 {
	return x * .1
}

func needSmall(x int) int { return x*10 + 1 }

func main() {
	fmt.Println(needSmall(Small))
	fmt.Println(needBig(Big))
}
