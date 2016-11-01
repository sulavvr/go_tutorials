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
		head := dl.first
		for head.Next != nil {
			head = head.Next
		}

		ins.Prev = head
		head.Next = ins
		dl.last = ins
	}
}

func (dl *DoublyLinkedList) Display() {
	head := dl.first

	for head != nil {
		fmt.Printf("%+v\n", head)
		head = head.Next
	}
}
