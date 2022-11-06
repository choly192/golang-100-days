package main
import (
	"fmt"
)
func main() {
	// 整型赋值
	var a =10
	b := a
	b = 101
	fmt.Printf("a：%v，a的内存地址是%p\n",a,&a)
	fmt.Printf("b：%v，b的内存地址是%p\n",b,&b)
	//数组赋值
	var c =[3]int{1,2,3}
	d := c
	d[1] = 100
	fmt.Printf("c：%v，c的内存地址是%p\n",c,&c)
	fmt.Printf("d：%v，d的内存地址是%p\n",d,&d)
}