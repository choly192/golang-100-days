package main
import (
	"fmt"
)

func main() {
	// 1. 直接声明
	var slice1 []int
	if slice1 == nil {
		fmt.Println("切片为空")
	}else{
		fmt.Println("切片不为空")
	}
	// 2. 声明并赋值
	var slice2 = []int{1,2,3}
	fmt.Println(slice2, len(slice2), cap(slice2)) // [1 2 3] 3 3
	// 3. 通过make函数声明  make([]type, len, cap)
	slice3 := make([]int, 2) // 省略容量 --- 表示切片的长度和容量都是 2, 初始值为 0
	slice4 := make([]string, 1 ,2)
	fmt.Println(slice3,slice3[0]) // [0 0] 0
	fmt.Println(slice4, len(slice4),cap(slice4)) // [] 1 2

	// 4. 从数组进行切片  区间-- 前包后不包
	arr := [5]int{1,2,3,4,5}
	slice5 := arr[:] // 切取整个数组
	fmt.Println(slice5) // [1 2 3 4 5]
	// 5. 从切片切取切片
	var slice6 = slice5[2:3]
	fmt.Println(slice6) // [3]
}