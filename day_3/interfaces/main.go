package interfaces

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type Copier interface {
	copy(source, dest string)
}

func printShape(s shape) {
	fmt.Printf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return r.width*2 + r.height*2
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 3.14 * 2 * c.radius
}

type copier struct {
}

func (c copier) copy(s string, d string) {
	fmt.Printf("Copied file from %v to %v\n", s, d)
}

func copy(c Copier, source string, dest string) {
	c.copy(source, dest)
}

func Interafaces() {
	c := circle{radius: 10}
	r := rect{width: 10, height: 4}
	printShape(c)
	printShape(r)
	copy(copier{}, "c", "d")
}
