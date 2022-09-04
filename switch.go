package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Printf("%s.\n", os)
	case "linux":
		fmt.Printf("%s.\n", os)
	default:
		fmt.Println("other system")
	}
}
