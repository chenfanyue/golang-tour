package main

import "fmt"

func main() {
	// 空接口类型可以保存任何底层类型
	var i interface{} = "hello"

	s := i.(string) // 断言值类型
	fmt.Println(s)

	s, ok := i.(string) // 一般用ok-idiom模式
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
