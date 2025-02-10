package maps

import "fmt"

func ShowNests() {
	people := map[string]map[string]string{
		"Alice": {
			"Age":  "25",
			"City": "New York",
		},
		"Bob": {
			"Age":  "23",
			"City": "Miami",
		},
	}

	fmt.Println("Total people:", len(people))
	fmt.Println("Alice ka samaan:", people["Alice"])
	fmt.Println("Bob's your uncle, & his city:", people["Bob"]["City"])
	fmt.Println("----------------------------")
}
