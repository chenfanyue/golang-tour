package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 1
	fmt.Println(m["Answer"])

	m["Answer"] = 2
	fmt.Println(m["Answer"])

	delete(m, "Answer")
	fmt.Println(m["Answer"])

	v, ok := m["Answer"]
	fmt.Println(v, ok)
}
