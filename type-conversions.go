package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var z float64 = float64(x*x + y*y)
	var f float64 = math.Sqrt(z)
	var u uint = uint(f)
	fmt.Printf("%T %T %T\n", z, f, u)
	fmt.Println(z, f, u)
}
