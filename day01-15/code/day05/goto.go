package main

import (
	"fmt"
)

func main() {
	var count int = 5

	LOOP: for count < 10 {
		if count == 8 {
			count += 1
			goto LOOP
		}
		fmt.Printf("count的值为 : %d\n", count)
		count++
		/*
		count的值为 : 5
		count的值为 : 6
		count的值为 : 7
		count的值为 : 9
		*/
	}

	fmt.Println("----------------")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakHere
			}
			fmt.Printf("i的值为%d j的值为%d\n",i,j)
		}
	}
	//手动返回，避免执行进入标签。。
	return

	breakHere: fmt.Println("已执行完毕。。。。")
	/*
	i的值为0 j的值为0
	i的值为0 j的值为1
	已执行完毕。。。。
	*/
}