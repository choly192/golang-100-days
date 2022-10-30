package main

import (
	"fmt"
)

func main() {
	var s string = "abc"
	// 可以使用特殊变量 _ 来忽略不要的变量值
	for i, v := range s {
		fmt.Println(i,string(v))
		/*
		0 a
		1 b
		2 c
		*/
	}
	num := [6]int{0,1,2,3}
	for i,v:=range num {
		fmt.Printf("第%d位的值是%d\n", i,v)
		/*
		第0位的值是0
		第1位的值是1
		第2位的值是2
		第3位的值是3
		第4位的值是0
		第5位的值是0
		*/
	}
}