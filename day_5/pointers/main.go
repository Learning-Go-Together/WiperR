package pointers

import "fmt"

func Pointers() {
	name := "Shivang"
	namePtr := &name
	*namePtr = "Dhruva"

	fmt.Println(name)
	*namePtr = "Akshat"
	fmt.Println(*namePtr)
}

func NullPointers() {
	name := "Nice"
	var s *string = &name
	if s != nil {
		fmt.Println(*s)
	}
}

type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func (c car) setColor2(color string) {
	c.color = color
}

func PointerReceiver() {
	c := car{color: "white"}
	c.setColor("red")
	fmt.Println(c.color)
	c.setColor2("white")
	fmt.Println(c.color)
}
