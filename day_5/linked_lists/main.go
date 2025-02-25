package LinkedList

import (
	"errors"
	"fmt"
)

type LinkedListBase struct {
	root    *Node
	current *Node
	length  int
}

func (l *LinkedListBase) Add(item int) {
	node := Node{item, nil}
	l.length++
	if l.current == nil {
		l.root = &node
		l.current = &node
	} else {
		l.current.next = &node
		l.current = &node
	}
}

func (l *LinkedListBase) Insert(idx, val int) error {
	if idx < 0 || idx > l.length {
		return errors.New("Length is out of bounds")
	}
	node := l.root
	cp := 0

	for cp < idx-1 {
		cp++
		node = node.next
	}

	newNode := Node{val, nil}
	if idx == 0 {
		l.root = &newNode
		newNode.next = node
	} else {
		temp := node.next
		node.next = &newNode
		newNode.next = temp
	}
	l.length++
	return nil
}

func (l LinkedListBase) Display() {
	temp := l.root

	for temp != nil {
		fmt.Printf("%v --> ", temp.val)
		temp = temp.next
	}
	fmt.Printf("nil\n")
}

func New() LinkedListBase {
	return LinkedListBase{nil, nil, 0}
}

type Node struct {
	val  int
	next *Node
}
