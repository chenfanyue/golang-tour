package main

import "fmt"

/*
 * fibonacci is a function that returns
 * a function that returns an int.
 */
func fibonacci() func() int {
	a, b, n := 0, 1, 0
	return func() int {
		if n < 2 {
			n++
			return n - 1
		} else {
			a, b = b, a+b
			n++
			return b
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
