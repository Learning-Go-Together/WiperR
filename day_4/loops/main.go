package loops

import "fmt"

func Loops() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func WhileLoopUsingFor() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
