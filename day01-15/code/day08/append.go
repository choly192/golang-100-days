package main

import (
	"fmt"
)

func main() {
	var slice1 = make([]string, 3, 5)
	slice1[0] = "张三"
	slice1[1] = "李四"
	slice1[2] = "王五"

	// 向切片追加 一个元素
	slice2 := append(slice1, "芭比Q")
	fmt.Println(slice1, slice2) // [张三 李四 王五] [张三 李四 王五 芭比Q]

	// 向切片追加多个元素
	slice3 := append(slice1, "test","test1")
	fmt.Println(slice3, len(slice3), cap(slice3)) // [张三 李四 王五 test test1] 5 5
	// 追加元素时 超过了 切片的容量 此时会重新分配底层数组，即使未填满
	slice4 := append(slice3,"test3")
	fmt.Println(slice4,len(slice4), cap(slice4)) // [张三 李四 王五 test test1 test3] 6 10

	// 切片拷贝
	var slice5 [] string = make([]string,6, 10)
	// 将 slice4的拷贝给 slice5
	copy(slice5, slice4)
	fmt.Println(slice5) // [张三 李四 王五 test test1 test3]
}