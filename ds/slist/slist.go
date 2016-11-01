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
		head := sl.first

		for head.Next != nil {
			head = head.Next
		}

		head.Next = ins
	}
}

func (sl *SinglyLinkedList) Display() {
	head := sl.first;
	for head != nil {
		fmt.Printf("%+v\n", head)
		head = head.Next
	}
}

func (sl *SinglyLinkedList) RemoveLastItem() {
	if sl.first == nil {
		fmt.Println(errors.New("Nothing to remove"))
	} else {
		head := sl.first
		prev := head
		for head.Next != nil {
			prev = head
			head = head.Next
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
