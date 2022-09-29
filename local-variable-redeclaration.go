package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		j := i + 1
		k := 0
		fmt.Printf("i, %v, %p\n", i, &i)
		fmt.Printf("j, %v, %p\n", j, &j)
		fmt.Printf("k, %v, %p\n", k, &k)
	}
}

/*
i, 0, 0xc00001c030
j, 1, 0xc00001c038
k, 0, 0xc00001c040
i, 1, 0xc00001c030
j, 2, 0xc00001c050
k, 0, 0xc00001c058
i, 2, 0xc00001c030
j, 3, 0xc00001c060
k, 0, 0xc00001c068
*/