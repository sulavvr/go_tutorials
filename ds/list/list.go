package list

import (
	"fmt"
	"errors"
	
	"github.com/user/ds/node"
)

type List struct {
    first *node.Node
}

func (l *List) Insert(insertData interface{}) {
    ins := &node.Node{Data: insertData}
    
    if l.first == nil {
        l.first = ins
    } else {
        x := l.first
        
        for x.Next != nil {
            x = x.Next
        }
        
        x.Next = ins
    }
}

func (l *List) Display() {
    x := l.first
    
    for {
        if x.Next != nil {
			fmt.Printf("%v(%T)\n", x, x.Data)
            x = x.Next
        } else {
			fmt.Printf("%v(%T)\n", x, x.Data)
            break
        }
    }
}

func (l *List) RemoveLastItem() {
	if l.first == nil {
		fmt.Println(errors.New("Nothing to remove"))
	} else {
		x := l.first
		prev := x
		for x.Next != nil {
			prev = x
			x = x.Next
		}
		
		prev.Next = nil
		
	}
	
	fmt.Println("Removed last item")
}