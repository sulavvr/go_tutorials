package main

import (
	"fmt"
	"bufio"
	"os"
	//"../module"
)

func main() {
	fmt.Println("Enter your name...");

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	fmt.Println("Hello", text)
}
