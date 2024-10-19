package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// this is my change
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	// no while loop, while loop in go implemented as a for loop

	value := 1
	for value < 10 {
		fmt.Println("Value: ", value)
		value++
	}
	// goto statement with labels
	// can also use continue or break
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum: ", sum)
		if sum > 200 {
			goto theEnd
		}
	}

	// label
theEnd:
	fmt.Println("End of program")
}
