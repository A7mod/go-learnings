package maps

import "fmt"

func CheckKeyExists() {
	scores := map[string]int{
		"Alice":  98,
		"Damian": 88,
	}

	// Checking if key exists
	value, exists := scores["Damian"]
	if exists {
		fmt.Println("Damian's Score:", value)
	} else {
		fmt.Println("Damian is non-existant")
	}

	// Checking a non-existent key
	value, exists = scores["Bob"]
	if exists {
		fmt.Println("Bob's Scores:", value)
	} else {
		fmt.Println("Bob is not here man. He doesn't exist!")
	}

	fmt.Println("----------------------------")
}
