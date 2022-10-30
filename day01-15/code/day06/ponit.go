package main

import (
	"fmt"
)

func main() {
	var num int = 10
	fmt.Printf("num的地址为：%p\n", &num)

	var x *int
	x = &num
	fmt.Println(x)
	fmt.Printf("x地址对应的值为：%d", *x)

	/*
	num的地址为：0xc0000a6058
	0xc0000a6058
	x地址对应的值为：10
	*/
}