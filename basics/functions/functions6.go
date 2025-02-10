package basics

import "fmt"

func Counter() func() int {
	count := 0          // Local variable inside the outer function.
	return func() int { // Inner Function (closure - jo mujhe kabhi mila nahi)
		count++
		return count
	}
}

// this will be called from main.go.   (Wah bhai main.go file toh Github p ebheji hi ni abhi. lowl.s)
func ShowClosures() {
	increment := Counter()
	fmt.Println("Count:", increment())
	fmt.Println("Count:", increment())
	fmt.Println("Count:", increment())
	fmt.Println("Count:", increment())
	fmt.Println("Count:", increment())
	fmt.Println("----------------------------")

}
