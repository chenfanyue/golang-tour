package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)
	
	lenOfItems(11, 22, 190)

	s = append(s, 2, 3, 4)
	printSlice(s)
	
	s = append(s, []int{5, 6}...)
	printSlice(s)
	
	for i, v := range s {
		fmt.Println(i, v)
	}
	
	for _, v := range s {
		fmt.Println(v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func lenOfItems(items ...int) {
	fmt.Println(len(items))
}
