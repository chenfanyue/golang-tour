package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, e T) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	i := []int{10, 20, 15, -10}
	fmt.Println(Index(i, 15))

	// Index also works on a slice of strings
	s := []string{"foo", "bar", "baz"}
	fmt.Println(Index(s, "hi"))
}
