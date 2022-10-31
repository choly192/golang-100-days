## 数组

**定义：**数组是同一种数据类型的固定长度的序列

- 一致性： 数组只能保存相同的数据类型元素，元素的数据类型可以是任何相同的数据类型
- 有序性： 数组中的元素是有序的，可以通过下标访问
- 不可变性： 数组一旦初始化，则长度（数组中的元素个数）不可改变

### 数组初始化

- 先声明在赋值
- 声明并赋值
- 不限长度 [...]
- 索引设置

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. 先声明在赋值
	var arr1 [5]int

	arr1[0] = 10
	arr1[1] = 11
	arr1[2] = 12

	fmt.Println(arr1) // [10 11 12 0 0] 未初始化的值表示为默认值

	// for _, v := range arr1 {
	// 	fmt.Println(v)
	// }
	// 2. 声明并赋值
	var arr2 = [5]string{"a","b","c","d","e"}
	fmt.Println(arr2)

	// 3. 不限长度 [...]
	var arr3 = [...]int{1,2,3}
	fmt.Println(arr3,reflect.TypeOf(arr3)) // [1 2 3] [3]int

	// 4. 索引设置
	var arr4 = [5]string{1:"a", 4: "d"}
	fmt.Println(arr4,reflect.TypeOf(arr4)) //[ a   d] [5]string
}
```

### 多维数组

```go
package main

import (
	"fmt"
	"reflect"
)
func main (){
    	// 多为数组的声明
	var arr5 = [3][4]int{{1,2,3,4},{7,8,9,10},{11,12,13,14}}
	fmt.Println(arr5) // [[1 2 3 4] [7 8 9 10] [11 12 13 14]]

	// 注意 最后的维度 不能[...], 如二维数组的第二维度不能[...]
	var arr6 = [...][3]int{{1,2,3}}
	fmt.Println(arr6,reflect.TypeOf(arr6)) // [[1 2 3]] [1][3]int
}
```

### 数组是值类型

go中的数组是值类型，而不是引用类型。这意味着当它们被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映。

### 通过 len(arr) 求取数组的长度