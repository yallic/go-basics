package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Lorem ipsum dolor sit amet"

	// Returns characters between 0 and 5
	str_1 := str[:5]

	// Returns last 4 characters
	str_2 := str[len(str)-4:]
	str_3 := fmt.Sprintf("%s ipsum dolor sit %s", str_1, str_2)

	fmt.Println(str_3)

	if strings.EqualFold(str_1, "LOrEM") {
		fmt.Println("str_1 equal to LorEM")
	}

	// control begginning character
	if strings.HasPrefix(str_1, "Lorem") {
		fmt.Println("lorem")
	}

	fmt.Println(strings.ToUpper(str_3))
}
