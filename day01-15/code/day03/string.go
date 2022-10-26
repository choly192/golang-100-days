package main

import (
	. "fmt"
	"strings"
)

func main() {
	// 1. 索引
	s := "hello go world!"

	Println(s[0]) // 104 --- 字符串存的是 ASCII码
	Println(string(s[0])) // h

	// 2. 字符串的切片 string[start, end)
	Println(s[0:4]) // hell
	Println(s[4:]) // o go world!
	Println(s[:]) // hello go world!

	// 3. 字符串拼接 +
	var s1 = "ping"
	var s2 = " ROG"
	s3 := s1 + s2
	Println(s3) // ping ROG

	// 4. 求字符串长度 len(string)
	l := len(s3)
	Println(l) // 8

	// 5. 字符串分割  strings.Split(s,sep string) []string
	var s4 = strings.Split(s1,"")
	var s5 = strings.Split(s1,"i")
	Println(s4) // [p i n g]
	Println(s5) // [p ng]

	// 6. 字符串 **join**， strings.Join(a[]string, sep string)
	s6 := strings.Join(s4,"")
	s7 := strings.Join(s4,"9")
	Println(s6) // ping
	Println(s7) // p9i9n9g
}