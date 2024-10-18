package main

import (
	"fmt"
	"sort"
)

func main() {
	// slices are resizable and can be sorted easily
	var colors = [3]string{"Red", "Green", "Blue"} // this is an array
	fmt.Println(colors)

	// this is an slice
	var colors2 = []string{"Red", "Green", "Blue"} // this is an array
	fmt.Println(colors2)
	colors2 = append(colors2, "Purple")
	fmt.Println(colors2)

	colors2 = append(colors2[1:len(colors2)])
	fmt.Println(colors2)

	colors2 = append(colors2[:len(colors2)-1])
	fmt.Println(colors2)

	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 43
	numbers[2] = 41
	numbers[3] = 1
	numbers[4] = 5
	fmt.Println(numbers)
	numbers = append(numbers, 60)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
