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