package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Printf("%T %g\n", time.Now(), time.Saturday)
	
	switch time.Friday {
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println(time.Friday - today)
	}
}
