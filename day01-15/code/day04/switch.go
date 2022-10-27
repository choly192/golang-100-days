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

	var score2 = 90
	switch {
		case score2>score:
			fmt.Println("score2更优秀")
		case score2<score:
			fmt.Println("score更优秀")
		case score2==score:
			fmt.Println("score2和score一样优秀")
		default:
			fmt.Println("默认")
	}

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