package main

import (
	"fmt"
	"time"
)

// func task(i int) {
// 	fmt.Println("Value is :", i)
// }

func main() {
	for i := 0; i < 10; i++ {
		// go task(i)

		go func(i int) {
			fmt.Println("Value is :", i)
		}(i)
	}

	time.Sleep(1 * time.Second)
}
