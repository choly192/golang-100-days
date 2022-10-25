## 基本语法

---

### 基本类型

- 整型

整型分为以下两个大类： 按长度分为：`int8`、`int16`、`int32`、`int64`对应的无符号整型：`uint8`、`uint16`、`uint32`、`uint64`

- 浮点型

Go 语言支持两种浮点型数：`float32`和`float64`。这两种浮点型数据格式遵循 IEEE 754 标准： float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：`math.MaxFloat32`。 float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：`math.MaxFloat64`。

- 复数

`complex64`和`complex128`

复数有实部和虚部，`complex64`的实部和虚部为 32 位，`complex128`的实部和虚部为 64 位。

- 布尔值

Go 语言中以`bool`类型进行声明布尔型数据，布尔型数据只有`true`和`false`两个值
**注意：**

> 布尔类型变量的默认值为 false。
> Go 语言中不允许将整型强制转换为布尔型.
> 布尔型无法参与数值运算，也无法与其他类型进行转换。

- 字符串

Go 语言里的字符串的内部实现使用 UTF-8 编码，单行字符转以双引号（英文）包裹，多行字符转以`反引号`包裹,默认值是** "" **

### 变量的声明

- go 语言中的变量必须要先声明后使用，同一作用域内不能重复声明变量。以 **var** 关键字声明变量, 以字母或下划线开头，由一个或多个字母、数字、下划线组成.

- 标准声明

```go
 // var 变量名 变量类型
 var count int
 var name string
 var isOk bool

 // 初始化值 ---- 不初始化值，变量的值为为 **默认值**
 count = 1
 name = "qs"
 isOk = true
```

---

- 批量申明

```go
var (
  count int
  name string
  isOk bool
)
```

- 类型推导

```go
// 声明变量并赋值，省略类型声明，会自行推导
var count = 1
var name = "test"
var isOk = true
```

- 简短声明

```go
// 申明并初始化值
count := 1
```

**这种简短声明只能在函数体内使用，不能用于声明全局变量**

- 匿名变量，匿名变量用于忽略某个值，用下划线\_表示

```go
package main
import "fmt"


func foo() (int, string){
	return 1,"test"
}

func main () {
  	value1,_ := foo()
	_,value2 := foo()
	fmt.Println("value1==",value1) // value1== 1
	fmt.Println("value2==",value2) // value2== "test"
}
```

**注意事项**
在同一个作用域中，已存在同名的变量，则之后的声明初始化，则退化为赋值操作---前提是，最少要有一个新的变量被定义，且在同一作用域

```go
package main

import (
	"fmt"
)

func main() {
	x := 140
	fmt.Println(&x)
	x, y := 200, "abc"
	fmt.Println(&x, x)
	fmt.Print(y)
}

/*
	0xc0000180a8
	0xc0000180a8 200
	abc
	*/

```

---

### 常量

- 常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值，用关键字`const`声明

- 常量可作为枚举，常量组，**常量组中如不指定类型和初始化值，则与上一行非空常量右值相同**

```go
package main

import (
	"fmt"
)


func main() {
	const (
		x uint32 = 12
		y
		z string = "abc"
		w
	)

	fmt.Printf("%T,%d\n",y,y)
	fmt.Printf("%T,%s", w,w)

	/*
	输出结果：
	uint32,12
  string,abc
	*/
}
```

- iota
  iota，特殊常量，可以认为是一个可以被编译器修改的常量

  - 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：

  ```go
  const (
    a = iota
    b
    c
  )
  ```

  - 用法

  ```go
  package main
  import "fmt"
  func main() {
    const (
    a = iota // 0
    b // 1 iota += 1
    c // 2 iota += 1
    d = "qw" // 独立值 iota += 1
    e = 100 //独立值 iota += 1
    f // 100 iota += 1
    g = "cc" // iota += 1
    h // cc iota += 1
    i // cc iota += 1
    j = iota //9 恢复计数
    k // 10
    )
    fmt.Println(a,b,c,d,e,f,g,h,i,j,k) // 0 1 2 qw 100 100 cc cc cc 9 10
  }
  ```

  如果中断 iota 自增，则必须显式恢复。且后续自增值按行序递增

  自增默认是 int 类型，可以自行进行显示指定类型

  数字常量不会分配存储空间，无须像变量那样通过内存寻址来取值，因此无法获取地址
