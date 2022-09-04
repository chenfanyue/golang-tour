package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hi"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5} // append 0 automatically
	fmt.Println(primes)
}
