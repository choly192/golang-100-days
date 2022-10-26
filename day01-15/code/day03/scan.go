package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	var x int
	var s string
	fmt.Println("请输入一个整数和一个字符串:")
	fmt.Scanln(&x,&s) // 读取键盘的输入，通过操作地址，赋值给x和y   阻塞式操作
	fmt.Printf("x的值:%d s的值:%s", x,s)

	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1)

}