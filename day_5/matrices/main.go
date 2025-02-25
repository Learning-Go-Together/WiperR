package matrices

import "fmt"

func PrintMatrix() {
	mat := [][]int{}

	for i := 0; i < 10; i++ {
		slice := []int{}
		for j := 0; j < 10; j++ {
			slice = append(slice, i*j)
		}
		mat = append(mat, slice)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(mat[i])
	}
}
