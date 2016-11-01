package main

import (
	"github.com/sulavv/ds/list"
	"github.com/sulavv/ds/queue"
	// "github.com/sulavv/ds/stack"
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
	queue := &queue.Queue{List: &list.List{}}

	queue.Enqueue("test")
	queue.Enqueue(1)
	queue.Enqueue("hello")
	queue.Enqueue([]int{1, 2, 3, 4})
	queue.Enqueue([]byte("hello"))
	queue.Enqueue(map[int]int{1: 1, 2: 3, 4: 5})
	queue.Enqueue([2]string{"test", "testing"})

	queue.Dequeue()

	queue.List.Display()
}
