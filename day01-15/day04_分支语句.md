## 分支语句

---

### 条件语句

- if 语句

  - 可省略条件表达式括号。
  - 持初始化语句，可定义代码块局部变量。
  - 代码块左 括号必须在条件表达式尾部。

  if 布尔表达式 {
  /_ 在布尔表达式为 true 时执行 _/
  }

```go
package main
import (
	"fmt"
)
func main() {
	score := 30
	var s string = "studygo"
	if score>10 {
		fmt.Println(s[0:5]) // study
	}

	if x:=10; score>x{
		fmt.Println(x) // 10
		x++
		fmt.Println(x) // 11
	}

	var count int = 5
	for i := 0; i < count; i++ {
		// fmt.Println(i) // 0 1 2 3 4
		if i == 1 {
			fmt.Printf("i的值：%d\n", i) //1
		} else {
			fmt.Println(i) // 0 2 3 4
		}
	}
}
```

### switch 语句

```
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

- 变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
  您可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。

```go
package main
import (
	"fmt"
)
func main() {
	var score int = 80
	std := "A"
	switch score {
		case 90:
			std = "A"
		case 80:
			std = "B"
		case 50,60,70:
			std ="C"
		default:
			std ="D"
	}

	switch {
	case std == "A":
		fmt.Println("优")
	case std == "B", std == "C":
		fmt.Println("良")
	case std == "D":
		fmt.Println("不及格")
	default:
		fmt.Println("不及格==")
	}
}
```

- 贯穿 fallthrough

```go
package main

import (
  "fmt"
)

func main() {
  	// 贯穿  fallthrough
	switch x:=10;x{
	default:
		fmt.Println(x)
	case 10:
		x += 15
		fmt.Println(x) // 25
		fallthrough
	case 11:
		x-=3
		fmt.Println(x) // 22
	}
}
```

_注意：_

- case 后的常量值不能重复
- case 后可以有多个常量值
- fallthrough 应该是某个 case 的最后一行。如果它出现在中间的某个地方，编译器就会抛出错误。

---
