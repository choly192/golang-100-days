## Go 语言环境搭建

下载安装：https://golang.google.cn/dl/ 或者 https://www.studygolang.com/dl
安装完成后使用 go version 查看是否安装成功检查安装版本，如果出现命令不识别之类的问题需要设置一下环境变量
敲黑板：运行命令`go env`

---

> GOROOT --- go 的安装路径

1. GOROOT 下的 bin 目录下的 go 文件是 go 的编译器，所有的命令都由它执行
2. GOROOT 下的 src 是 go 的标准库，我们可以直接用的
   > GOPATH --- Go 语言中使用的一个环境变量，它使用绝对路径提供项目的工作目录（workspce）
3. GOPATH 下创建 src 文件夹，即存放 Go 项目代码的位置。
4. GOPATH 需要在 "我的电脑" -> "属性" -> "高级系统设置" -> "环境变量" 设置 GOPATH 的路径
5. GOPATH 对应创建的文件夹里面，手动创建 binpkgsrc 三个文件夹
   a. src --- 存储 go 的源代码（需要手动创建），之后你所有的代码都放在 src 文件夹中。
   b. pkg --- 存储编译后生成的包文件（自动生成）
   c. bin --- 存储生成的可执行文件（自动生成）
   > GOPROXY --- go 的代理，从 github 上拉取代码很慢或者拉不下来，需要设置代理
6. 命令 `go env -w GOPROXY=https://goproxy.cn,direct`对代理进行修改。
7. 阿里云：export GOPROXY=https://mirrors.aliyun.com/goproxy/
8. 七牛云：export GOPROXY=https://goproxy.cn
   > GO111MODULE --- 是否启用 go module 功能（模块支持功能）
9. `go env -w GO111MODULE=on`--- 开启模块支持，无论什么情况，go 命令都会使用 module
10. `go env -w GO111MODULE=off`--- 关闭模块支持，无论什么情况，go 命令都不会使用 module
11. `go env -w GO111MODULE=auto`---- 即当前目录在 GOPATH/src 之外且该目录包含 go.mod 文件时开启 module 功能；否则继续使用 GOPATH
    <b>注：可以先关闭 go module 避免理解混淆（学习包管理时选择）<b>

---

## 第一个 go 程序

- 创建 helloworld.go 文件

```go
package main // 申明main包 表示当前是一个可执行程序

import "fmt" // 导入 fmt这个标准苦=库

func main() { // main函数  程序的入口
	fmt.Println("hello world! go") // 终端打印
}
```

- 代码执行
  > 方法一: go run 命令
- 进入终端后，cd 到 helloworld.go 目录下 执行 `go run helloworld.go` 观察终端执行结果

- 或者 如这样执行 `go run .\day01-15\code\day01\helloworld.go`

  > 方法二：go build 命令

- 在 day01 目录下执行 `go build`,然后会生成一个 `day01.exe`的可执行文件

- 在 day01 目录下执行`./day01`,观察程序执行结果

---

## import 导入库

- <b>点操作</b>

```go
import (
  . "fmt"
)
```

这个点操作的含义就是这个包导入之后在你调用这个包的函数时, 你可以省略前缀的包名, 也就是前面你调
用的`fmt.Println("hello world")`可以省略的写成`Println("hello world")`.

- <b>别名</b>：我们可以把包命名成另一个我们用起来容易记忆的名字

```go
import (
  f "fmt"
)
```

如`f.Println("hello world")`.

- <b> \_ 操作</b>: 操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的 init 函数

```go
import (
  _ "github.com/gin-gonic/gin"
)
```

## go 常用命令

- `go env`用于打印 Go 语言的环境信息。

- `go run`命令可以编译并运行命令源码文件。

- `go get`可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装。

- `go build`命令用于编译我们指定的源码文件或代码包以及它们的依赖包。

- `go install`用于编译并安装指定的代码包及它们的依赖包。

- `go clean`命令会删除掉执行其它命令时产生的一些文件和目录。

- `go doc`命令可以打印附于 Go 语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。

- `go test`命令用于对 Go 语言编写的程序进行测试。

- `go list`命令的作用是列出指定的代码包的信息。

- `go fix`会把指定代码包的所有 Go 语言源码文件中的旧版本代码修正为新版本的代码。

- `go vet`是一个用于检查 Go 语言源码中静态错误的简单工具。

- `go tool pprof`命令来交互式的访问概要文件的内容。

---

## 编码规范

### 命名规范

- go 在命名时以字母或者下划线开头，go 在命名时不允许使用特殊符号，如 ‘@，%，$’, go 语言是一门对大小写敏感的强类型语言。

> 1. 当命名以一个大写字母开头(常量、变量、类型、函数名、结构字段等等)，如：Gin，那么<b>使用这种形式的标识符的对象就可以被外部包的代码所使用（先导包）</b>，这称为导出（类似 public 关键字）
> 2. <b>命名如果是以小写字母开头的，则表明对包外是闭合不可见的，但是在整个包的内部是可见并可用的</b>(类似 private 关键字)

#### package 命名规范

- 保持 package 的名字和目录一致，包名应该是简单而有意的命名，不要和标准库冲突。包名使用小写字母，不使用混合大小写或者下划线.

```go
package main
package test
```

- 文件命名：使用简短而有意义的命名，使用小写字母，单词之间以下划线分割. 如： first_pro.go

#### 结构体 命名规范

- 采用驼峰命名法，首字母大小写根据是否允许外部访问控制

- struct 声明和初始化采用多行格式，如：

```go
// 声明
type LoginInfo struct {
  name  string
  pwd   string
}
// 初始化
l := LoginInfo {
  name:  "choly192",
  pwd:   "123456"
}
```

#### 接口 命名规范

- 采用驼峰命名法，首字母大小写根据是否允许外部访问控制

- 单个函数的命名以 "er" 结尾, 如：Tester;

```go
type Tester interface {
    Test(p []byte) (n int, err error)
}
```

#### 变量 命名规范

- 变量名称一般遵循驼峰法，首字母根据访问控制原则大写或者小写,但遇到特有名词时，需要遵循以下规则
  - 如果变量为私有，且特有名词为首个单词，则使用小写，如: httpRequest;
  - 其它情况都应当使用该名词原有的写法,如: HTTPRequest;
  - 若变量类型为 bool 类型，则名称应以 Has, Is, Can 或 Allow 开头
  ```go
  var isExist bool
  var hasConflict bool
  var canManage bool
  var allowGitHook bool
  ```

#### 常量 命名规范

- 常量使用全部使用大写字母，以下划线分割单词

```go
const COUNT_SUM = 1
```

- 枚举类型的常量，需要先创建相应的类型

```go
type Scheme string

const (
    HTTP  Scheme = "http"
    HTTPS Scheme = "https"
)
```

### 注释

- 单行注释 "//"
- 多行注释也叫块注释，均已以 /_ 开头，并以 _/ 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段;
- 注释风格 参考 https://golang.google.cn/doc/comment

### 其余规范遇到补充
