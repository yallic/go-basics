package main

import "fmt"

// We can store data in key value format with maps in go

//When we define a map we write key type first and then we write value type
var map_1 map[int]string

func main() {

	// We must initialize defined map that out of main function
	// If we not define map in the main func  will throw panic
	// Because map have not a reference address
	map_1 = make(map[int]string)
	map_2 := make(map[int]int)
	map_3 := make(map[int]int, 3)

	// We can give map lenght with 3rd parameter like in the above example

	map_1[0] = "1"
	map_2[0] = 2
	map_3[0] = 3
	map_3[1] = 3
	map_3[2] = 3
	map_3[3] = 3

	fmt.Println(map_1, map_2, map_3)
	fmt.Printf("Lenght of map_3: %d \n", len(map_3))
}
