package main

import "fmt"

func main() {
	var (
		i int
		f float64
		b bool
		s string
	)
	
	fmt.Printf("%v %g %q\n", i, i, i)
	fmt.Printf("%v %v %q", f, b, s)
}

// 0 %!g(int=0) '\x00'
// 0 false ""
