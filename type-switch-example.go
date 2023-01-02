package main

import (
	"fmt"
)

func main() {

	var t interface{}
	var i = 1
	t = &i
	switch t := t.(type) {
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	}
}
