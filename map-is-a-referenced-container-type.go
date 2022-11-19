package main

import "fmt"

func add(m map[int]int) {
	fmt.Printf("%p\n", m)
	m[2] = 1001
	fmt.Printf("%p, %#v\n", m, m)
}

func main() { // map是引用类型
	a := map[int]int{0: 199, 1: 900}
	fmt.Printf("%p, %#v\n", a, a)
	add(a)
}
/*
0xc00010c180, map[int]int{0:199, 1:900}
0xc00010c180
0xc00010c180, map[int]int{0:199, 1:900, 2:1001}
*/
