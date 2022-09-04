package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m == nil)
	
	m["Bell Labs"] = 11
	fmt.Println(m["Bell Labs"])
}
