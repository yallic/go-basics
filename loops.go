package main

import (
	"fmt"
	"time"
)

func main() {

	slc_1 := []int{1, 2, 3}

	for i, value := range slc_1 {
		fmt.Printf("index: %d value: %d \n", i, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d \n", i)
	}

	// for creating a endless while loop we can with for {}
	// endless loop

	c := time.After(5 * time.Second)
	for {

		b := true

		select {
		case <-c:
			b = true
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}

		if b {
			break
		}

	}
}
