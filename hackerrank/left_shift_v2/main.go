package main

/**
 * Hackerrank
 * Cracking the coding interview > Arrays: Left Rotation
 */
import "fmt"

func main() {
	var (
		n int
		d int
	)

	fmt.Scanf("%d %d", &n, &d)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scanf("%d", &x)

		arr[(i+n-d)%n] = x
	}

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}

}
