package main

/**
 * Hackerrank
 * Algorithm > Warmup > Time Conversion
 */
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var time string

	fmt.Scanf("%s", &time)

	ap := string(time[8:])
	time = string(time[0:8])

	x := strings.Split(time, ":")

	conv, _ := strconv.Atoi(x[0])

	twef := 12

	if ap == "AM" {
		twef = conv
		if conv == 12 {
			twef = 0
		}
	} else {
		if conv != 12 {
			twef = conv + 12
		}
	}

	var y string
	if twef < 10 {
		y = "0" + strconv.Itoa(twef)
	} else {
		y = strconv.Itoa(twef)
	}

	fmt.Printf("%s:%s:%s", y, x[1], x[2])

}
