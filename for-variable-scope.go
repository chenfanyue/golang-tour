package main

import (
	"fmt"
	"unsafe"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Print(i, unsafe.Pointer(&i), " ")
		i++
		i := i // 这里声明的变量i遮挡了上面声明的i, 右边的i为上面声明的循环变量i
		fmt.Print(i, unsafe.Pointer(&i), " ")
		i = 10 // 新声明的i被更改了
		fmt.Print(i, " ")
		_ = i
		fmt.Println(i)
	}
}

/*
0 0xc0000b2000 1 0xc0000b2020 10 10
2 0xc0000b2000 3 0xc0000b2028 10 10
*/
