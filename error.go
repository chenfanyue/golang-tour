package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"an error message",
	}
	// return &MyError{
	// 	When: time.Now(),
	// 	What: "an error message",
	// }
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		fmt.Printf("%T", err)
	}
}
