package range_go

import "fmt"

func Range() {
	nums := []int{1, 2, 3, 4}
	for i, num := range nums {
		fmt.Println(i, num)
	}
}
