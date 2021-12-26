package main

import "fmt"

// slices is like arrays but they dont need to be defined lenght of array

var slc_1 []int

func main() {

	// when slices to be defined with make func we need to give two parameter
	// actually we can do definition without second parameter
	// first parameter is used for slices length
	// second paramete is capacity of slice . capacity is max length of slice  but it can be changable
	slc_2 := make([]int, 0, 3)
	fmt.Printf("slc_2 len:%d cap:%d \n", len(slc_2), cap(slc_2))
	// In the below declaration will give to an error bcs slc_2's length is 0

	//slc_2[0] = 2

	// But this state is valid for assignment with "=" character
	// If we use append function, we will not get any error
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 2)
	slc_2 = append(slc_2, 3)

	fmt.Printf("slc_1 len:%d cap:%d \n", len(slc_1), cap(slc_1))
	fmt.Printf("slc_2 len:%d cap:%d \n", len(slc_2), cap(slc_2))

	// If we append to new value to slc_2, slc_2's capacity will double
	// In the golang slices logic like this

	slc_2 = append(slc_2, 4)

	fmt.Printf("slc_2 len:%d cap:%d \n", len(slc_2), cap(slc_2))

	// If we dont give any capacity value when defining slice , it takes deafult capacity . Default capacity is 2 in the golang
	slc_3 = make([]int, 0)
	fmt.Printf("slc_3 len:%d cap:%d \n", len(slc_3), cap(slc_3))

}
