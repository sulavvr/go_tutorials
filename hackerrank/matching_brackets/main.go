package main

/**
 * Hackerrank
 * Cracking the coding interview > Stacks: Balanced Brackets
 */

import (
	"fmt"
)

var stack []string

var brackets = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func main() {

	var count int
	var strings []string

	fmt.Scanf("%d", &count)

	for i := 0; i < count; i++ {
		var str string
		fmt.Scanf("%s", &str)

		strings = append(strings, str)
	}

	// {[()]}

	// [ {, [, (,  ]
	for _, v := range strings {
		for _, x := range v {
			if _, ok := brackets[string(x)]; ok {
				push(string(x))
			} else {
				if exists := checkIfPrev(string(x)); exists {
					pop()
				}
			}
		}
		if len(stack) > 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
		clearStack()
	}

}

func checkIfPrev(val string) bool {
	if len(stack) > 0 {
		in_stack := stack[len(stack)-1] // (
		if v, _ := brackets[in_stack]; v == val {
			return true
		}
	}
	return false
}

func push(val string) {
	stack = append(stack, val)
}

func pop() {
	stack = stack[:(len(stack) - 1)]
}

func clearStack() {
	stack = []string{}
}
