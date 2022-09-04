package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	fmt.Println(m == nil) // true
	m = make(map[string]Vertex)
	fmt.Println(m == nil) // false
	
	m["Bell Labs"] = Vertex{
		-1.1, 2.2, // Lat: -1.1, Long: 2.2,
	}
	fmt.Println(m["Bell Labs"])
}
