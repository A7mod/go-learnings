package structs

import "fmt"

func ShowAnonStruct() {
	//  Defining and using an anonymous struct (no explicit type definition)
	book := struct {
		Title  string
		Author string
	}{
		Title:  "Normal People",
		Author: "Sally Rooney",
	}

	// Used for temporary struct definitions (when you don’t need to reuse the type).
	// Best for short-lived data structures that don’t need to be reused elsewhere.

	// Printing the struct values
	fmt.Println("Book Information Ferda:", book)
	fmt.Println("----------------------------")

}
