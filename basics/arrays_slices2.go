package basics

import "fmt"

func ShowSlices() {

	slice := []int{1, 2, 3, 4, 5} // slice declaration (this isan array tho!)
	fmt.Println("Slice:", slice)

	// append elements in that array.
	slice = append(slice, 6, 7, 8)
	fmt.Println("After Append:", slice)

	arr := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println("Array before:", arr)
	slice2 := arr[1:4]
	slice3 := slice2[1:5] // interesting revelation this slice3 still pulls items from arr and not slice3.
	fmt.Println("Sliced Array slice2:", slice2)
	fmt.Println("Sliced Array slice3:", slice3)
}
