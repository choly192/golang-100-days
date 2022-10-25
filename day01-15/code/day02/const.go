package main

import (
	"fmt"
)


func main() {
	const (
		x uint32 = 12
		y
		z string = "abc"
		w
	)

	fmt.Printf("%T,%d\n",y,y)
	fmt.Printf("%T,%s", w,w)

	/*
	输出结果：
	uint32,12
  string,abc
	*/
}