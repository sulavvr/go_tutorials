package dlist

import (
	// "errors"
	"fmt"

	"github.com/sulavvr/ds/node"
)

type DoublyLinkedList struct {
	first *node.Node
	last  *node.Node
}

func (dl *DoublyLinkedList) GetFirst() *node.Node {
	return dl.first
}

func (dl *DoublyLinkedList) GetLast() *node.Node {
	return dl.last
}

func (dl *DoublyLinkedList) Insert(data interface{}) {
	ins := &node.Node{Data: data}
	if dl.first == nil {
		dl.first = ins
		dl.last = ins
	} else {
		x := dl.first
		for x.Next != nil {
			x = x.Next
		}

		ins.Prev = x
		x.Next = ins
		dl.last = ins
	}
}

func (dl *DoublyLinkedList) Display() {
	x := dl.first

	for {
		if x.Next != nil {
			fmt.Printf("(%p)%+v\n", x, x)
			x = x.Next
		} else {
			fmt.Printf("(%p)%+v\n", x, x)
			break
		}
	}
}
