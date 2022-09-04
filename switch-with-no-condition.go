package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	hour := t.Hour()
	switch {
	case hour < 12:
		fmt.Println("mornig")
	case hour < 15:
		fmt.Println("afternoon")
	default:
		fmt.Println("evenig")
	}
}

func main() {
	t := time.Now()
	hour := t.Hour()
	if hour < 12 {
		fmt.Println("mornig")
	} else if hour < 15 {
		fmt.Println("afternoon")
	} else {
		fmt.Println("evenig")
	}
}
