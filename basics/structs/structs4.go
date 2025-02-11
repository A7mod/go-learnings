package structs

import "fmt"

type Address struct {
	City, State string
}

// user structs contains an address field
type User struct {
	Name    string
	Age     int
	Address Address // Embedding Address Struct here
}

func ShowUser() {
	// Creating an instance for user and initializing Address
	user := User{
		Name: "Nush",
		Age:  23,
		Address: Address{
			City:  "New York City",
			State: "NY",
		},
	}
	fmt.Println("User Address:", user)
	fmt.Println("User:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("City:", user.Address.City)
	fmt.Println("State:", user.Address.State)

	fmt.Println("----------------------------")
}
