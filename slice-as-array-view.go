package main

import "fmt"

func main() {
	names := [4]string{
		"a0",
		"b1",
		"c2",
		"d3",
	}
	fmt.Println(names)
	
	a := names[:2]
	b := names[1:3]
	fmt.Println(a, b)
	
	b[0] = "other"
	fmt.Println(names)
	fmt.Println(a, b)
}
