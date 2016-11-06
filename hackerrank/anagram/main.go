package main

/**
 * Hackerrank
 * Cracking the coding interview > Strings: Making Anagrams
 */
import (
	"fmt"
)

func main() {
	var a, b string
	var count int

	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)

	arr_a := make(map[string]int)

	for _, v := range a {
		if val, ok := arr_a[string(v)]; ok {
			arr_a[string(v)] = (val + 1)
		} else {
			arr_a[string(v)] = 1
		}
	}

	for _, v := range b {
		if val, ok := arr_a[string(v)]; ok {
			if val > 1 {
				arr_a[string(v)] = (val - 1)
			} else {
				delete(arr_a, string(v))
			}
		} else {
			count++
		}
	}

	var arr_count int
	for _, v := range arr_a {
		arr_count += v
	}

	fmt.Printf("%d", count+arr_count)

}
