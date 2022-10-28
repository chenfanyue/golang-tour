package main

import "fmt"

type OrderStatus int

const (
	NoPay OrderStatus = iota
	Pending
	Delivered
	Received
)

func main() {
	fmt.Printf("%#v %p\n", Pending, Pending)
	fmt.Printf("%#v %p\n", Received, Received)
}

