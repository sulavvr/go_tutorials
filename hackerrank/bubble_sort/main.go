package main

/**
 * Hackerrank
 * Cracking the coding interview > Sorting: Bubble Sort
 */
import (
	"fmt"
)

func main() {
	var count int
	var arr []int
	var num int

	fmt.Scanf("%d", &count)

	for i := 0; i < count; i++ {
		var x int
		fmt.Scanf("%d", &x)
		arr = append(arr, x)
	}

	for i := 0; i < count; i++ {
		for j := 0; j < count-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				num++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", num)
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[count-1])
}
