package main

import "fmt"

const Pi = 3.14

func main() {
	const Hi = "世界"
	const Truth = true
	
	fmt.Printf("type %T\n", Pi)
	fmt.Println("say hi", Hi)
	fmt.Println("it's", Truth)
}
