package main

import "fmt"

func main() {
	var arr [5]int // Fixed size array (of 5 integers)
	// usually arrays have all values default 0
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	fmt.Println("Array:", arr)
	fmt.Println("Element at index 1:", arr[1])

	// array with values
	numbers := [4]int{100, 200, 300} // this is also one way to add items in the array, nice!
	fmt.Println("Numbers:", numbers)

}
