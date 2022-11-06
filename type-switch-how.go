package main

import "fmt"

type Vertex struct {
	X int
	Y float64
}

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("its int\n")
	case string:
		fmt.Printf("str\n")
	default:
		fmt.Printf("%T, ", i)
		fmt.Println(i)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	do([]int{11, 22})
	do(Vertex{900, 000})
}
