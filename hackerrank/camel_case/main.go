package main

/**
 * Hackerrank
 * Algorithms > Strings > CamelCase
 */

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Scanf("%s", &str)

	count := 1
	for _, v := range str {
		if strings.ToLower(string(v)) != string(v) {
			count++
		}
	}

	fmt.Println(count)
}
