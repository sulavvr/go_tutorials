package main

/*
 * Hackerrank
 * Data structures > Arrays > Sparse Arrays
 */

import (
	"fmt"
)

func main() {
	var (
		num_strings int
		strings     []string
		num_queries int
		q_strings   []string
		// count_arr   []int
	)

	fmt.Scanf("%d", &num_strings)

	for i := 0; i < num_strings; i++ {
		var str string
		fmt.Scanf("%s", &str)
		strings = append(strings, str)
	}

	fmt.Scanf("%d", &num_queries)

	for i := 0; i < num_queries; i++ {
		var str string
		fmt.Scanf("%s", &str)
		q_strings = append(q_strings, str)
	}

	for _, v := range q_strings {
		count := 0
		for _, val := range strings {
			if v == val {
				count++
			}
		}
		// count_arr = append(count_arr, count)
		fmt.Println(count)
	}

	// for _, v := range count_arr {
	// 	fmt.Printf("%d\n", v)
	// }

}
