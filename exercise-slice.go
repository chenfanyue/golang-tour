package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for i := range arr {
		arr[i] = make([]uint8, dx)
	}
	return arr
}

func main() {
	pic.Show(Pic)
}
