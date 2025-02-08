package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6)
	fmt.Println("Appended slice (slice2):", slice2)

	// copying slices
	original := []int{10, 20, 30}
	copied := make([]int, len(original))
	copy(copied, original)
	fmt.Println("Copied Slice:", copied)

	// length and capacity
	fmt.Println("Length 1:", len(slice1), "Capacity 1:", cap(slice1))
	fmt.Println("Length 2:", len(slice2), "Capacity 2:", cap(slice2))
	fmt.Println("Length original:", len(original), "Capacity original:", cap(original))
	fmt.Println("Length copied:", len(copied), "Capacity copied:", cap(copied))

	// this was useless

}
