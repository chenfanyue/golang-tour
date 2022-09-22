package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	fmt.Printf("%v %T\n", r, r)

	b := make([]byte, 8)
	fmt.Printf("%v %T\n", b, b)
	for {
		n, err := r.Read(b)
		fmt.Printf("%v %T\n", r, r)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
