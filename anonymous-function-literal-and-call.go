package main

import "fmt"

func main() {
	func(str string) {
		fmt.Println(str)
	}("hello, golang")
}
