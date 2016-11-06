package main

/**
 * Hackkerrank
 * Cracking the coding interview > Hash Tables: Ransom Note
 */
import "fmt"

func main() {
	var (
		mag_count int
		ran_count int
		ret_val   string
	)

	mag_words := make(map[string]int)

	fmt.Scanf("%d %d", &mag_count, &ran_count)

	for i := 0; i < mag_count; i++ {
		var x string
		fmt.Scanf("%s", &x)
		if v, ok := mag_words[x]; ok {
			mag_words[x] = (v + 1)
		} else {
			mag_words[x] = 1
		}
	}

	ret_val = "Yes"

	for i := 0; i < ran_count; i++ {
		var x string
		fmt.Scanf("%s", &x)
		if v, ok := mag_words[x]; !ok {
			ret_val = "No"
			break
		} else {
			if v > 1 {
				mag_words[x] = (v - 1)
			} else {
				delete(mag_words, x)
			}
		}
	}

	fmt.Printf("%s", ret_val)
}
