package slices

import "fmt"

func Slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:len(primes)])
}
