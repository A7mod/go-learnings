package maps

import "fmt"

func ModifyMaps() {
	student := map[string]int{
		"Alice": 78,
		"Bob":   68,
		"John":  93,
	}
	fmt.Println("Initial map:", student)

	// adding a new key-value pair
	student["Eve"] = 92
	fmt.Println("After adding Eve:", student)

	// Updating an existing value.
	student["Bob"] = 88
	fmt.Println("After updating item for Bob", student)

	// Deleting a key
	delete(student, "John")
	fmt.Println("After the delete operation:", student)
	fmt.Println("----------------------------")
}
