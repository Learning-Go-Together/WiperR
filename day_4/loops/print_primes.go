package loops

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	fmt.Printf("Printing prime numbers upto %v\n", max)
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
	fmt.Println("===========================================")
}

func PrimePrintMain() {
	printPrimes(10)
	printPrimes(20)
	printPrimes(30)
}
