package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "javagood"
	ch <- "phpnotbad"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
