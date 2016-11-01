package slist

import (
	"errors"
	"fmt"

	"github.com/sulavvr/ds/node"
)

type SinglyLinkedList struct {
	first *node.Node
}

func (sl *SinglyLinkedList) Insert(insertData interface{}) {
	ins := &node.Node{Data: insertData}

	if sl.first == nil {
		sl.first = ins
	} else {
		x := sl.first

		for x.Next != nil {
			x = x.Next
		}

		x.Next = ins
	}
}

func (sl *SinglyLinkedList) Display() {
	x := sl.first

	for {
		if x.Next != nil {
			fmt.Printf("%+v\n", x)
			x = x.Next
		} else {
			fmt.Printf("%+v\n", x)
			break
		}
	}
}

func (sl *SinglyLinkedList) RemoveLastItem() {
	if sl.first == nil {
		fmt.Println(errors.New("Nothing to remove"))
	} else {
		x := sl.first
		prev := x
		for x.Next != nil {
			prev = x
			x = x.Next
		}

		prev.Next = nil

	}

	fmt.Println("Removed last item")
}

func (sl *SinglyLinkedList) RemoveFirstItem() {
	if sl.first == nil {
		fmt.Println(errors.New("Nothing to remove"))
	} else {
		sl.first = sl.first.Next
	}
}
