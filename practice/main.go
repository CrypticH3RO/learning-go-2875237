package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	// like in java, characters are wrapped in single quotes while
	// while full strings are wrapped in double quotes
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
}
