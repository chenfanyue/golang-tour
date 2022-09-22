package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (list *List[T]) traverse() {
	for list != nil {
		fmt.Print(list.val)
		list = list.next
	}
}

func main() {
	a := List[int]{next: nil, val: 33}
	b := List[int]{&a, 22}
	c := List[int]{&b, 11}
	c.traverse()
}
