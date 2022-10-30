package main

import (
	"fmt"
)


func main() {
	// 1. 基本形式
	for i := 0; i < 5; i++ {
		fmt.Println(i) // 0 1 2 3 4
	}
	// 2. for condition {}
	var s string = "abtest"
	l := len(s)
	for l > 0 { 
		fmt.Println(l)
		l--
	}
}