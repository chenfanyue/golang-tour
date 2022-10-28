package main

import "fmt"

func main() {
	const (
		a int = iota
		b
		c
		d float64 = 11
		e
		f         = iota
		g
	)
	fmt.Printf("%#v %p\n", a, a)
	fmt.Printf("%#v %p\n", b, b)
	fmt.Printf("%#v %p\n", c, c)
	fmt.Printf("%#v %p\n", d, d)
	fmt.Printf("%#v %p\n", e, e)
	fmt.Printf("%#v %p\n", f, f)
	fmt.Printf("%#v %p\n", g, g)
}
