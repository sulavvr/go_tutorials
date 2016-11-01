package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Guess a number between 1 to 20")

	var guess int
	_, err := fmt.Scan(&guess)

	if err != nil {
		fmt.Printf("Huge mistake - %s", err)
	}
	r := rand.New(rand.NewSource(99))

	randomNum := r.Intn(20)

	if guess < randomNum {
		fmt.Printf("Your guess is lower than the number. The number was %v\n", randomNum)
	} else if guess > randomNum {
		fmt.Println("Your guess is higher than the number. The number was %v\n", randomNum)
	} else {
		fmt.Println("Whoaa! You guessed the right number")
	}

}
