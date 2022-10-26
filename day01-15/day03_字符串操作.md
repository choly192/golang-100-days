## 字符串操作

---

### 字符串操作

- 字符串的 切片和索引操作  
  go 语言中字符串的索引吗，没有*负索引*, 索引从 **0** 开始
  切片操作区间是*左闭右开*, 字符串的切片 string[start, end)

- 字符串拼接 **+**

- 求字符串的长度 **len(string)**

- 字符串分割 strings.Split(s,sep string) []string

  - s -> 要分割的字符串
  - sep -> 要去掉的字符串

- 字符串 **join**， strings.Join(a[]string, sep string)
  - s -> 可用来连接元素的字符串
  - sep -> 放置在最终字符串中元素之间的分隔符

```go
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
```

- 字符串更多操作可看 **strings**包中的方法

---

## 格式化打印

- %v,原样输出
- %T，打印类型
- %t,bool 类型
- %s，字符串
- %f，浮点
- %d，10 进制的整数
- %b，2 进制的整数
- %o，8 进制
- %x，%X，16 进制
- %x：0-9，a-f
- %X：0-9，A-F
- %c，打印字符
- %p，打印地址

---

## 键盘输入

### fmt 包读取键盘输入

```go
package main

import (
	"fmt"
)

func main() {

	var x int
	var s string
	fmt.Println("请输入一个整数和一个字符串:")
	fmt.Scanln(&x,&s) // 读取键盘的输入，通过操作地址，赋值给x和y   阻塞式操作
	fmt.Printf("x的值:%d s的值:%s", x,s)

}
```

### bufio 包读取

```go
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1)

}
```
