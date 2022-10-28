package main

import "fmt"

func main() {
	a := []int{10, 11, 12, 13, 14}
	b := a[2:]
	fmt.Printf("%#v %p\n", a, a)
	fmt.Printf("%#v %p\n", b, b)

	n := copy(a, b)
	fmt.Printf("%#v %p\n", a, a)
	fmt.Printf("%#v %p\n", b, b)
	fmt.Println(n)
}

/*
[]int{10, 11, 12, 13, 14} 0xc000100000
[]int{12, 13, 14} 0xc000100010
[]int{12, 13, 14, 13, 14} 0xc000100000
[]int{14, 13, 14} 0xc000100010
3
*/
