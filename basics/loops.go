package basics

import "fmt"

func ShowLoops() {
	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Count:", i)
	}

	// For range
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println("Index:", index, "value:", value)
	}
}
