package main

import "fmt"

func main() {
	ret, _ := f1("a")
	fmt.Println(ret)
}
func f1(in string) (string, string) {
	return "ahgp", in + "la"
}
