package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	// his version doesn't wokr, chatGPT said some servers block request without a User-Agent
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Print(content)

	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}

}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20) // initalize size as 0 and initial capacity at 20, this is a slice
	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	// reads next avialalbe object from the json object
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}

/*
chatGpt version which works
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	// Create a new request and add headers (User-Agent)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	// Setting a User-Agent header
	req.Header.Set("User-Agent", "Go-HTTP-Client")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response type: %T\n", resp)

	// Check if the response status is 200 OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	// Read the response body
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)

	// Check if the content is JSON before decoding
	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
		fmt.Println("Error: Response is not JSON")
		fmt.Println(content)
		return
	}

	fmt.Print(content)

	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20) // initialize size as 0 and initial capacity at 20
	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	// reads next available object from the JSON array
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
*/
