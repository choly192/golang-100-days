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