package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		poodle := Dog{"Poodle", 10, "Woof!"}
		fmt.Println(poodle)
		fmt.Printf("%+v\n", poodle)
		fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

		poodle.Speak()
		poodle.Sound = "Arf!"
		poodle.Speak()
		poodle.SpeakThreeTimes()
		poodle.SpeakThreeTimes()
	*/
	// file system part
	content := "Hello from Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()                 // defer waits till eveyrthing else is done befroe closing the file
	defer readFile("./fromString.txt") // defer wait until the file completely closeed until you read it.
}

func readFile(fileName string) {
	// data, err := ioutil.ReadFile(fileName) // this is depecrated
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file: ", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// SpeakThreeTimes is how the dog speaks loudly
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
