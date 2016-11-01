package main

import (
	// "fmt"
	// "github.com/sulavvr/ds/queue"
	// "github.com/sulavvr/ds/slist"
	"github.com/sulavvr/ds/dlist"
	// "github.com/sulavvr/ds/stack"
)

func main() {
	// Stack
	// stack := &stack.Stack{List: &list.List{}}

	// stack.Push("test")
	// stack.Push(1)
	// stack.Push("hello")
	// stack.Push([]int{1, 2, 3, 4})
	// stack.Push([]byte("hello"))
	// stack.Push(map[int]int{1:1, 2:3, 4:5})
	// stack.Push([2]string{"test", "testing"})

	// stack.Pop()

	// Queue
	// queue := &queue.Queue{List: &slist.SinglyLinkedList{}}

	// queue.Enqueue("test")
	// queue.Enqueue(1)
	// queue.Enqueue("hello")
	// queue.Enqueue([]int{1, 2, 3, 4})
	// queue.Enqueue([]byte("hello"))
	// queue.Enqueue(map[int]int{1: 1, 2: 3, 4: 5})
	// queue.Enqueue([2]string{"test", "testing"})

	// queue.Dequeue()

	// queue.List.Display()
	//
	list := &dlist.DoublyLinkedList{}
	list.Insert("test")
	list.Insert(1)
	list.Insert("hello")
	list.Insert([]int{1, 2, 3, 4})
	list.Insert([]byte("hello"))
	list.Insert(map[int]int{1: 1, 2: 3, 4: 5})
	list.Insert([2]string{"test", "testing"})
	//
	list.Display()

	// fmt.Println(list.GetFirst(), list.GetLast())

}
