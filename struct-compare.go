package main

import (
	"fmt"
	//"reflect"
)

type User struct {
	Id   int
	Name string
}

func main() {
	var u1 User
	u2 := User{}
	compareObj(u1, u2)
	//fmt.Printf("%+v\n", u)
	//fmt.Println(reflect.TypeOf(struct{}{}))
	//fmt.Printf("%v \n", &struct{}{})
}

func compareObj(o1, o2 interface{}) {
	fmt.Printf("o1: %#v\n", o1)
	fmt.Printf("o2: %#v\n", o2)
	fmt.Printf("o1-p: %p\n", o1)
	fmt.Printf("o2-p: %p\n", o2)
	fmt.Printf("&o1: %p\n", &o1)
	fmt.Printf("&o2: %p\n", &o2)
	fmt.Printf("o1 == o2: %t\n", o1 == o2)
	fmt.Printf("&o1 == &o2: %t\n", &o1 == &o2)
}
