package structs

import "fmt"

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) diameter() int {
	return r.height*2 + r.width*2
}

func StructMethods() {
	r := rect{width: 10, height: 4}
	fmt.Println(r.area())
	fmt.Println(r.diameter())
}
