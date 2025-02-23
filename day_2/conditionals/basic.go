package conditionals

import "fmt"

func Basic() {
	if age := 18; age >= 18 {
		fmt.Println("You are old enough")
	} else {
		fmt.Println("You are not old enough")
	}
}
