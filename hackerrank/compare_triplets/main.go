package main

/*
 * Hackerrank
 * Algorithms > Warmup > Compare the triplets
 */
import (
	"fmt"
)

func main() {
	var (
		alice_score []int
		bob_score   []int
		alice_point int
		bob_point   int
	)

	for j := 0; j < 3; j++ {
		var x int
		fmt.Scanf("%d", &x)
		alice_score = append(alice_score, x)
	}

	for i := 0; i < 3; i++ {
		var x int
		fmt.Scanf("%d", &x)
		bob_score = append(bob_score, x)
	}

	for i := 0; i < 3; i++ {
		if alice_score[i] > bob_score[i] {
			alice_point++
		} else if alice_score[i] < bob_score[i] {
			bob_point++
		}
	}

	fmt.Printf("%d %d", alice_point, bob_point)

}
