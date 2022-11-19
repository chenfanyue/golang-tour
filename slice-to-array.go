package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	p := (*[3]int)(s)
	a := *p
	fmt.Println(a)
	// 切片和指针是引用类型，数组是值类型
	fmt.Printf("%p\n%p\n%p\n", s, p, &a)
}
/*
[1 2 3]
0xc0000b2000
0xc0000b2000
0xc0000b2018
*/
