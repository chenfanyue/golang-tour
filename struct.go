package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Printf("%T", Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 9
	fmt.Println(v)

	p := &v
	p.X = 1e9
	// (*p).X = 1e9
	fmt.Println(v)
	fmt.Printf("%T", p)
}
