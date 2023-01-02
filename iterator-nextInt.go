package main

import (
	"fmt"
)

func main() {
	var t, next = nextInt([]byte("ab123cd321efg"), 1)
	fmt.Println(t, next)
}

func nextInt(b []byte, i int) (int, int) {
	isDigit := func(ch byte) bool {
		return '0' <= ch && ch <= '9'
	}
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}
