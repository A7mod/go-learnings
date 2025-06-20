package con2

import (
	"fmt"
	"time"
)

func TickBasic() {
	ticker := time.Tick(1 * time.Second)

	for i := 0; i < 5; i++ {
		t := <-ticker
		fmt.Println("Tick at:", t)
	}
}
