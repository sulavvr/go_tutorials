package main

/**
 * Hackerrank
 * Algorithms > Implementation > Sock Merchant
 */
import "fmt"

func main() {
	var total int

	fmt.Scanf("%d", &total)

	var count int

	arr := make(map[int]int)

	for i := 0; i < total; i++ {
		var x int
		fmt.Scanf("%d", &x)
		if _, ok := arr[x]; ok {
			delete(arr, x)
			count++
		} else {
			arr[x] = 0
		}
	}

	fmt.Println(count)
}
