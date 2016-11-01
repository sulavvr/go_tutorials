package queue

import (
	"github.com/sulavvr/ds/slist"
)

type Queue struct {
	List *slist.SinglyLinkedList
}

func (q *Queue) Enqueue(data interface{}) {
	q.List.Insert(data)
}

func (q *Queue) Dequeue() {
	q.List.RemoveFirstItem()
}
