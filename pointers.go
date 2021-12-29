package main

import (
	"fmt"
	"strings"
)

//  nil
type String *string

func main() {

	var rstr String
	var pstr string

	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "fatih"
	rstr = &pstr

	//adres
	fmt.Println(rstr)
	//adresdeki değer
	fmt.Println(*rstr)
	fmt.Println(pstr)

	pstr = "pstr ile degistirdim"
	fmt.Println(*rstr)
	fmt.Println(pstr)

	*rstr = "rstr ile degistirdim"
	fmt.Println(*rstr)
	fmt.Println(pstr)

	replaceStr(rstr)

	fmt.Println(*rstr)
	fmt.Println(pstr)
}
func replaceStr(s String) {
	*s = strings.Replace(*s, "rstr", "RSTR", 1)
}
