package main

import (
	"fmt"
)

func main() {
	doSomething()
	sum := addValues(5, 8)
	fmt.Println(sum)

	multisum, multiCount := addAllValues(4, 7, 9)
	fmt.Println(multisum)
	fmt.Println(multiCount)

	multisum2, multiCount2 := addAllValues(4, 7)
	fmt.Println(multisum2)
	fmt.Println(multiCount2)
}

func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(a int, b int) int {
	return a + b
}

// can do this if same type
func addValues2(a, b int) int {
	return a + b
}

// can have as many values as we want
func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
