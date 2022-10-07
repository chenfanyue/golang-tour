package main

import "fmt"

func do[T int|bool](in T) {
	fmt.Println(in)
}

func main() {
	do(101)
	do(1 > 0)
}
