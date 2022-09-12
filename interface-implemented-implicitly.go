package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
	i int
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello", 10}
	i.M() // i只能访问接口I定义过的方法
	fmt.Println(i.i) // i.i undefined (type I has no field or method i)
}
