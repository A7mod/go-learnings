package maps

import "fmt"

func IterateMaps() {
	countries := map[string]string{
		"USA": "United States of America",
		"IN":  "India",
		"UK":  "United Kingdom",
		"AUS": "Australia",
	}

	fmt.Println("Country Names ? :")
	for key, value := range countries {
		fmt.Printf("%s -> %s \n", key, value)
	}
	fmt.Println("----------------------------")
}
