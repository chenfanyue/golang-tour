package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("10,")
	if err != nil {
		fmt.Printf("%v, %T", err, err)
		return
	}
	fmt.Printf("%v, %T", i, i) // err != nil时i为0值
}
