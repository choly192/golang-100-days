## 切片
- `切片` 是数组的一个引用，是引用类型，但其本身是一个结构体，值拷贝传递
- 切片的是一个长度可变的的数组
- 切片的初始值是 `nil`, 长度`len`和`cap`容量都是 `0`

---

### 切片的构成
1. 指针，指向数组中切片指定的起始位置
2. 长度, 即切片的长度
3. 容量, 切片的容量，切片的容量总是大于等于切片的长度的

- 切片的数据结构定义：
```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```

### 切片的创建
1. 直接声明, 语法 `var identifier []type`
2. 声明并赋初值
3. 通过make函数声明切片, 语法 `make([]type, len, cap)`
4. 从数组切片
5. 从切片进行切片 

```go
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
```

### append函数 和 copy函数

--- 
1、每次 append 操作都会检查 slice 是否有足够的容量，如果足够会直接在原始数组上追加元素并返回一个新的 slice，底层数组不变，但是这种情况非常危险，极度容易产生 bug！而若容量不够，会创建一个新的容量足够的底层数组，先将之前数组的元素复制过来，再将新元素追加到后面，然后返回新的 slice，底层数组改变而这里对新数组的进行扩容

2、扩容策略：如果切片的容量小于 1024 个元素，于是扩容的时候就翻倍增加容量。上面那个例子也验证了这一情况，总容量从原来的4个翻倍到现在的8个。一旦元素个数超过 1024 个元素，那么增长因子就变成 1.25 ，即每次增加原来容量的四分之一。

- 基本使用

```go
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
```

### 切片的插入和删除

- 在切片的开头插入元素

```go
package main

import (
	"fmt"
)

func main() {
	var slice [] int = []int{1,2,3}
	fmt.Println(slice) // [1 2 3]
	// 1. 在前面插入一个切片, 后面的切片需要 解包
	slice = append([]int{9,8,7},slice...)
	fmt.Println(slice) // [1 2 3]
}
```
> 在切片开头添加元素一般都会导致内存的重新分配，而且会导致已有元素全部被复制 1 次，因此，从切片的开头添加元素的性能要比从尾部追加元素的性能差很多。

- 在切片的任意位置插入元素

```go
var a []int
a = append(a[:i], append([]int{x}, a[i:]...)...) // 在第i个位置插入x
a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片
```

- 删除元素

  + go语言中没有专门删除切片元素的方法，可以利用切片本身的特性来删除元素
  + `要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)`


```go
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
```

### 切片元素排序

```go
package main
import (
	"sort"
	"fmt"
)

func main() {
	var slice []int = []int{6,2,9,100}
	sort.Ints(slice)
	fmt.Println(slice) // [2 6 9 100]
	b:=[]string{"melon","banana","caomei","apple"}

	sort.Strings(b)
	fmt.Println(b) // [apple banana caomei melon]

	c:=[]float64{3.14,5.25,1.12,4,78}
	sort.Float64s(c)
	fmt.Println(c) // [1.12 3.14 4 5.25 78]

	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	sort.Sort(sort.Reverse(sort.Float64Slice(c)))
	fmt.Println(slice, c) // [100 9 6 2] [78 5.25 4 3.14 1.12]
}
```
