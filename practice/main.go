package main

import (
	"fmt"
)

func main() {

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[0])

	var nums = [5]int{5, 4, 3, 2, 1}
	fmt.Println(nums)

	fmt.Println("length of colors", len(colors))
}
