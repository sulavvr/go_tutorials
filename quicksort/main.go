package main

import (
	"fmt"
)

func main() {
	arr := []int{35, 33, 42, 10, 14, 19, 27, 44, 26, 31}

	fmt.Println(partition(arr))
}

func partition(arr []int) []int {
	arr_len := len(arr)
	pivot := arr[arr_len-1]
	lo, hi := 0, arr_len-2

	for {
		for arr[lo] < pivot {
			lo++
		}

		for arr[hi] > pivot && hi > 0 {
			hi--
		}

		if lo >= hi {
			break
		} else {
			arr[lo], arr[hi] = arr[hi], arr[lo]
		}
	}

	arr[lo], arr[arr_len-1] = arr[arr_len-1], arr[lo]

	return arr
}
