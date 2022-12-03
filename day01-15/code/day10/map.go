package main

import "fmt"

func main() {
	// 1. 直接声明赋值
    info := map[string]string{"name":"张三","age":"24"}
    fmt.Println(info) // map[age:24 name:张三]
	
	// 2. 先声明再赋值
	var stus = make(map[string]string);
	stus["name"] = "李思思"
	fmt.Println(stus) // map[name:李思思]

	var countryCapitalMap map[string]string // nil
	/* 创建集合 */
	countryCapitalMap = make(map[string]string) // 必须要开辟存储空间后才能操作
	countryCapitalMap["城市"] = "成都"
	fmt.Println(countryCapitalMap) // map[城市:成都]

}