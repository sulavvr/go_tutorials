package queue

import (
	"github.com/sulavv/ds/list"
)

type Queue struct {
	List *list.List
}

func (q *Queue) Enqueue(data interface{}) {
	q.List.Insert(data)
}

func (q *Queue) Dequeue() {
	q.List.RemoveFirstItem()
}
