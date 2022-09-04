package main

import (
	"fmt"
)

func main() {
	switch s := "abcd"; s[1] {
	case 'a':
		fmt.Println('a')
		fallthrough
	case 'b':
		fmt.Printf("%T\n", s[1])
		fmt.Println('b', "b")
		fallthrough
	case 'c':
		fmt.Println('c')
		fallthrough
	case 'd':
		fmt.Println('d' > 'a')
	default:
		fmt.Println("other character")
	}
}
