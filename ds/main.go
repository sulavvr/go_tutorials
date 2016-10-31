package main

import (
    "github.com/user/ds/list"
    "github.com/user/ds/stack"
)

func main() {
    stack := &stack.Stack{List: &list.List{}}
    
    stack.Push("test")
    stack.Push(1)
    stack.Push("hello")
    stack.Push([]int{1, 2, 3, 4})
    stack.Push([]byte("hello"))
    stack.Push(map[int]int{1:1, 2:3, 4:5})
    stack.Push([2]string{"test", "testing"})
    
//     stack.Pop()
    
    
    stack.List.Display()
}