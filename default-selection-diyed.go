package main

import (
	"fmt"
	"time"
)

func do(c chan string) {
	c <- "datashit"
}

func main() {
	t := time.Now()
	var i int
	d := make(chan string)
	go do(d)
	for i < 4 {
		select {
		case s := <-d:
			fmt.Println(i, time.Since(t), s)
			i++
		default:
			fmt.Println(i, time.Since(t), "--default--")
			i++
		}
	}
}
