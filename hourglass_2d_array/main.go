package main

import (
	"fmt"
)

func main() {
	var arr [][]int

	for i := 0; i < 6; i++ {
		var x []int
		for j := 0; j < 6; j++ {
			var in int
			fmt.Scan(&in)
			x = append(x, in)
		}
		arr = append(arr, x)
	}

	//t.Println(arr)
	max := -99

	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[i])-2; j++ {
			new_sum := sums(arr[i][j:j+3], arr[i+1][j+1], arr[i+2][j:j+3])

			if max < new_sum {
				max = new_sum
			}
		}
	}

	fmt.Println(max)
}

func sums(x []int, y int, z []int) int {
	var s int
	for _, v := range x {
		s += v
	}
	for _, v := range z {
		s += v
	}

	s += y

	return s
}
