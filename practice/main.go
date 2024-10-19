package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	// no break command after each case, automatically goes to the end when one condition is true
	switch dow {
	case 1:
		result = "It's Sunday!"
		// fallthrough // this makes it so ti checks the other cases still
	case 2:
		result = "It's Monday!"
		// fallthrough
	default:
		result = "It's some other day!"
	}
	fmt.Println(result)

	switch dow2 := rand.Intn(7) + 1; dow2 {
	case 1:
		result = "It's Sunday!"
		// fallthrough // this makes it so ti checks the other cases still
	case 2:
		result = "It's Monday!"
		// fallthrough
	default:
		result = "It's some other day!"
	}
	fmt.Println(result)
}
