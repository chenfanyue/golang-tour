package main

import "fmt"

func main() {
	i, j := 1, 9
	p := &i
	
	*p = 2
	fmt.Println(i)
	
	p = &j
	*p /= 3
	fmt.Println(j)

	k := p
	fmt.Printf("%T\n", k)
	*k *= 5
	fmt.Println(j)
}

func main() {
	i, j := 150, 9
	fmt.Printf("%T\n", i)

	p := &i         // point to i
	fmt.Printf("%T\n", p)
	fmt.Println(*p) // read i through the pointer
	*p = 11         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p /= 3   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
