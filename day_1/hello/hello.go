package hello

import (
	"example.com/greetings"
	"example.com/hbday"
	"fmt"
)

func Hello() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)

	message = hbday.Hbday("Gladys")
	fmt.Println(message)

	message = hbday.HbdayBad("Gladys")
	fmt.Println(message)
}
