package node

import (
    "fmt"
)

type Node struct {
    Data interface{}
    Next *Node
}

func (n *Node) GetValue() interface{} {
    return fmt.Sprintf("%v", n.Data)
}