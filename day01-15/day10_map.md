## map

Map是一种无序的键值对的集合。Map最重要的一点是通过key来快速检索数据。可以像数组和切片的迭代一样来迭代它。Map是无序的，我们无法决定它的返回顺序，因为它是使用hash表来实现的。Map是引用类型。

---
*注意：*
- map是无序的。每次打印的map会不一样，不能通过index获取，必须通过key值获取对应的值。
- map的长度是不固定的，也就是和slice一样，是引用类型。
- 内置的`len`函数同样适用于map，返回map拥有的key的数量。

> slice查询是遍历方式，时间复杂度是O(n), map查询是hash映射 ;当数据量小的时候切片查询比map快，但是数据量大的时候map的优势就体现出来了

---

### map的声明和初始化

声明语法
```go
// map_name为变量名  key_type为键类型 key_value为键对应的值类型
var map_name map[key_type]value_type
```
1. 直接声明赋值
```go
package main

import "fmt"

func main() {
    info := map[string]string{"name":"张三","age":"24"}
    fmt.Println(info)
}
```
2. 通过make函数声明
```go
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
```
*注意：*
```go
/* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type
```

### map的删除 - delete函数

delete 函数用于删除集合的元素, 参数为 map 和其对应的 key。删除函数不返回任何值。

```go
delete(map,key)
```

### 判断map中某个键是否存在

```go
value,ok := map[key] // 如果key存在，ok-> true; value是key对应的值

```