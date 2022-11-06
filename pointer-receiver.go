package main

import "fmt"

type user struct {
	id   int
	name string
	age  uint16
}

func (u *user) info() string {
	fmt.Printf("%p\n", u)
	str := fmt.Sprintf("I am %s, My id is %d, and my age is %d.\n", u.name, u.id, u.age)
	return str
}

func main() {
	u := user{
		id:   1,
		name: "lucy",
		age:  20,
	}
	fmt.Printf("%p\n", &u)
	userInfo := u.info()
	fmt.Println(userInfo)
}
