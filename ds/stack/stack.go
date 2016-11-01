package stack

import (
	"github.com/sulavvr/ds/slist"
)

type Stack struct {
	List *slist.SinglyLinkedList
}

func (s *Stack) Push(data interface{}) {
	s.List.Insert(data)
}

func (s *Stack) Pop() {
	s.List.RemoveLastItem()
}
