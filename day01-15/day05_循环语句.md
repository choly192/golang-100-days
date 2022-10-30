## 循环语句
go 语言中只有 **for** 循环这唯一一种循环结构，，没有**while**循环语句

### for 循环的形式
```
    for init; condition; post { }
    for condition { }
    for { }
    init： 一般为赋值表达式，给控制变量赋初值；
    condition： 关系表达式或逻辑表达式，循环控制条件；
    post： 一般为赋值表达式，给控制变量增量或减量。
```

```go
package main

import (
	"fmt"
)


func main() {
	// 1. 基本形式
	for i := 0; i < 5; i++ {
		fmt.Println(i) // 0 1 2 3 4
	}
	// 2. for condition {}
	var s string = "abtest"
	l := len(s)
	for l > 0 { 
		fmt.Println(l)
		l--
	}
}
```

### for 循环的变体---range
for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环

- 格式
```
for key, value := range oldMap {
    newMap[key] = value
}
```

```go
package main

import (
	"fmt"
)

func main() {
	var s string = "abc"
	// 可以使用特殊变量 _ 来忽略不要的变量值
	for i, v := range s {
		fmt.Println(i,string(v))
		/*
		0 a
		1 b
		2 c
		*/
	}
	num := [6]int{0,1,2,3}
	for i,v:=range num {
		fmt.Printf("第%d位的值是%d\n", i,v)
		/*
		第0位的值是0
		第1位的值是1
		第2位的值是2
		第3位的值是3
		第4位的值是0
		第5位的值是0
		*/
	}
}
```

### 循环控制语句 break continue goto
- break 终止循环并跳出
- continue 跳出当前循环，不影响下一循环的执行
- goto 可以无条件地转移到过程中指定的行。

```go
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
```