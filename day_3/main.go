package main

import (
	// "day_3/billing_system"
	"day_3/closures_adder"
	"fmt"
)

func main() {
	addr := closure_sadder.Addr()
	fmt.Println(addr(1))
	fmt.Println(addr(6))
	fmt.Println(addr(8))
}
