package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["wa"] = "washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"
	fmt.Println(states)
	// order is not guarenteed for maps
	for k, v := range states {
		fmt.Printf("%v: %v\n ", k, v)
	}

	// to print in order, convert to a slice, sort slice of keys
	// then use the slice to output values in the map
	keys := make([]string, len(states))

	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
