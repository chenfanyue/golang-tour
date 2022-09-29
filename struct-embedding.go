package main

import "fmt"

type base struct {
	color string
}

func (b *base) say() {
	fmt.Println("Hi from say function", b)
}

type child struct {
	base  //embedding
	style string
}

func main() {
	/*
	child := &child{
		base: base{color: "Red"},
		style:"somestyle",
	} */
	child := &child{
		base{color: "Red"},
		"somestyle",
	}
	child.say()
	fmt.Println("The color is " + child.color)
}
