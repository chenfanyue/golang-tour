package main

import (
	"fmt"
)

func main() {
	var c = make(chan int)
	go func(d chan int) {
		d <- 123
	}(c)
	fmt.Println(c)
	fmt.Printf("%p\n", c)
	fmt.Println(<-c)
}
