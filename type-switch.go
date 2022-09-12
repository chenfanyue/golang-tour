package main

import "fmt"

type Vertex struct {
	X int
	Y float64
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v * 2 = %v\n", v, v*2)
	case string:
		fmt.Printf("%q, %v bytes\n", v, len(v))
	default:
		fmt.Printf("%T, ", v)
		fmt.Println(v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	do([]int{11, 22})
	do(&Vertex{900, 000})
}
