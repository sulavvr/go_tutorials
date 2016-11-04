package main

/**
 * Hackerrank
 * Algorithm > Warmup > Plus Minus
 */

import (
	"fmt"
)

func main() {
	var (
		total int
		zeros int
		negs  int
		pos   int
	)

	fmt.Scanf("%d", &total)

	for i := 0; i < total; i++ {
		var x int
		fmt.Scanf("%d", &x)
		if x < 0 {
			negs++
		} else if x > 0 {
			pos++
		} else {
			zeros++
		}
	}

	fmt.Printf("%f\n%f\n%f", float64(pos)/float64(total), float64(negs)/float64(total), float64(zeros)/float64(total))
}
