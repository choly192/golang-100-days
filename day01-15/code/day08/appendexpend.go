package main

import (
	"fmt"
)

func main() {
	var slice [] int = []int{1,2,3}
	fmt.Println(slice) // [1 2 3]
	// 1. 在前面插入一个切片, 后面的切片需要 解包
	slice = append([]int{9,8,7},slice...) // 底层数组已经发生变化
	fmt.Println(slice) // [9 8 7 1 2 3]

	// 2. 在切片 slice中，索引为2的元素前插入一个元素或者切片
	slice = append(slice[:2],append([]int{100}, slice[2:]...)...)
	fmt.Println(slice) // [9 8 100 7 1 2 3]

	// 3. 从切片中删除元素   （删除 索引为4的元素）
	slice = append(slice[:4],slice[5:]...)
	fmt.Println(slice) // [9 8 100 7 2 3]
}