package con2

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroups() {
	var wg sync.WaitGroup

	wg.Add(3) // starting 3 Goroutines

	go func() {
		defer wg.Done() //notify completion
		time.Sleep(1 * time.Second)
		fmt.Println("🍞 Bread is toasted")
	}()

	go func() {
		defer wg.Done() //notify completion
		time.Sleep(2 * time.Second)
		fmt.Println("🥚 Eggs are boiled")
	}()

	go func() {
		defer wg.Done() //notify completion
		time.Sleep(3 * time.Second)
		fmt.Println("☕ Tea is ready")
	}()

	// wait until all goroutines finish!!

	wg.Wait()
	fmt.Println("Breakfast is ready!!")

}
