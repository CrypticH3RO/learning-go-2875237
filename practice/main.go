package main

import (
	"fmt"
)

func main() {
	anInt := 42

	var p *int // default value is nil
	var p2 = &anInt
	fmt.Println("Value of p:", p)
	fmt.Println("Value of p2:", *p2)

	value1 := 42.13
	pointer1 := &value1 // pointing at the memory address not the value

	fmt.Println("Value 1: ", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer 1: ", *pointer1)
	fmt.Println("Value 1: ", value1)

}
