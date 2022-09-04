package main

import "fmt"

func main() {
	fmt.Println("push into defer stack")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("push done")
}
