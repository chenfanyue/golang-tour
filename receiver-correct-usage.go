package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	fmt.Printf("ptr %p\n", r)
	return r.width * r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Printf("ori %p\n", &r)

	fmt.Println("area: ", r.area())

	rp := &r
	fmt.Println("area: ", rp.area())
}
