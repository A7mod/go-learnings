package basics

import "fmt"

func ShowMaps() {
	// Declaring an empty map
	var person map[string]int
	fmt.Println("Empty map:", person) // empty map output

	// Initializing a map using make()
	person = make(map[string]int)
	person["age"] = 25
	fmt.Println("Person age:", person["age"])

	countryCodes := map[string]string{
		"US": "United States",
		"IN": "India",
		"UK": "United Kingdom",
	}
	fmt.Println("Country Codes:", countryCodes)
}
