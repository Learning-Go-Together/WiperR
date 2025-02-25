package main

import (
	"day_5/linked_lists"
)

func main() {
	root := LinkedList.New()
	root.Add(1)
	root.Add(2)
	root.Add(3)
	root.Add(4)
	root.Display()

	root.Insert(0, 5)
	root.Display()
	root.Insert(2, 6)
	root.Display()
}
