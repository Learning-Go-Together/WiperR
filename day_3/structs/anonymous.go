package structs

import "fmt"

func AnonymousStruct() {
	car := struct {
		brand string
		model string
	}{
		model: "Nothing",
		brand: "Shivang",
	}
	fmt.Println(car)
}
