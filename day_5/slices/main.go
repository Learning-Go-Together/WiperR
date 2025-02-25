package slices

import "fmt"

func Slices() {
	s := []int{1, 2, 3, 4}
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
}
