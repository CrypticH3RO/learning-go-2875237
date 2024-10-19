package main

import (
	"fmt"
)

func main() {
	poodle := Cat{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("The breed is %v\nWeight:%v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 9
	fmt.Printf("The breed is %v\nWeight:%v\n", poodle.Breed, poodle.Weight)
}

// upper case means exported , else not expored

/*
Just like the name of the type, the property names are either exported or non exported.
And again, it's determined by the initial character. Uppercase means that it's exported,
and the rest of the application can read and write to those property values.
*/
// Cat is a struct
type Cat struct {
	Breed  string
	Weight int
}
