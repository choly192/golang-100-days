package main

import (
	"fmt"
)

func main() {
	const (
		a = iota // 0
		b // 1 iota += 1
		c // 2 iota += 1
		d = "qw" // 独立值 iota += 1
		e = 100 //独立值 iota += 1
		f  // 100 iota += 1
		g = "cc" // iota += 1
		h // cc iota += 1
		i // cc iota += 1
		j = iota //9 恢复计数
		k // 10
	)
	fmt.Println(a,b,c,d,e,f,g,h,i,j,k) // 0 1 2 qw 100 100 cc cc cc 9 10
}