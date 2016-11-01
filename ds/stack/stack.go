package stack

import (
	"github.com/sulavv/ds/list"
)

type Stack struct {
	List *list.List
}

func (s *Stack) Push(data interface{}) {
	s.List.Insert(data)
}

func (s *Stack) Pop() {
	s.List.RemoveLastItem()
}
