package main

import "fmt"

func split(sum int) (x, y int) {
	// go的下整除往0方向靠拢
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(25))
}
