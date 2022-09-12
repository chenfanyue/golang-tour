package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 引用传递的方式访问，直接原地操作
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.X, v.Y)
}
/* 值传递的方式访问，在副本上操作
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.X, v.Y)
}
*/

func main() {
	v := Vertex{3, 4}
	fmt.Printf("%T\n", v)
	v.Scale(10)
	fmt.Println(v, v.Abs())
}
