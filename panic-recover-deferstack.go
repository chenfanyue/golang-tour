package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()

	defer func() {
		panic("exit2...")
	}()

	panic("exit...")
	fmt.Println("game over!")
}
