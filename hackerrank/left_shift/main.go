package main

/*
 * Hackerrank
 * Data structures > Arrays > Left Rotation
 */

import (
	"fmt"
)

func main() {
	var (
		count int
		shift int
		arr   []int
		temp  int
	)

	fmt.Scan(&count, &shift)
	for i := 0; i < count; i++ {
		var x int
		fmt.Scan(&x)
		arr = append(arr, x)
	}

	for i, j := 0, 0; i < shift; i++ {
		y := 0
		temp = arr[j]
		for y < len(arr)-1 {
			arr[y] = arr[y+1]
			y++
		}
		arr[len(arr)-1] = temp

	}

	for _, x := range arr {
		fmt.Printf("%d ", x)
	}
}
