package main
import "fmt"


func foo() (int, string){
	return 1,"test"
}

func main() {
	// 1. 标准声明
	var count int
	var isOk bool

	// fmt.Printf("count的type: %T\n, 数值是: %d", count,count)
// 变量初始化值
	count = 1
	isOk = true
	fmt.Printf("count的type: %T\n, 数值是: %d\n", count,count) // count的type: int, 数值是: 1

	// 写在一行
	var name string = "test"
	fmt.Printf("name的type: %T\n, 数值是: %s\n", name,name)  // name的type: string, 数值是: test
	fmt.Printf("isOk的type: %T\n, 数值是: %v\n", isOk,isOk)  // isOk的type: bool, 数值是: true

	// 2. 类型推导
	var stu = "test"
	fmt.Printf("stu的类型是：%T\n stu的值是：%s\n", stu,stu)  // stu的类型是：string stu的值是：test

	// 3. 简短声明
	shot := "qs"
	fmt.Printf("shot的类型是：%T\n shot的值是：%s\n", shot, shot) // shot的类型是：string shot的值是：qs
	
	// 4. 批量申明
	var (
		a = 3
		b = "李四"
		isShow = false
	)
	fmt.Printf("a：%d b: %s isShow: %v\n", a,b,isShow) // a：3 b: 李四 isShow: false

	// 5. 匿名变量，匿名变量用于忽略某个值，用下划线_表示
	value1,_ := foo()
	_,value2 := foo()
	fmt.Println("value1==",value1) // value1== 1
	fmt.Println("value2==",value2) // value2== "test"

	// 注意  在同一个作用域中，已存在同名的变量，则之后的声明初始化，则退化为赋值操作---前提是，最少要有一个新的变量被定义，且在同一作用域
	x := 140
	fmt.Println(&x)
	x, y := 200, "abc"
	fmt.Println(&x, x)
	fmt.Print(y)
	/*
	0xc0000180a8
	0xc0000180a8 200
	abc
	*/
} 