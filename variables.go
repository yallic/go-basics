package main

import "fmt"

var degisken_1 string
var degisken_2 = "deger 2"

func main() {

	degisken_1 = "deger 1"
	degisken_3 := "deger 3"
	// Go da variable tanımı yaparken 3 farklı opsiyon var

	// 1- var keyword ve tipi belirleyerek
	// 2- var keyword ve tip belirlemeyerek
	// 3- := işareti ile

	// := işareti func çağırıları dışında yapılmaz çünkü func dışında tüm ifadeler keyword ile başlamalıdır.

	fmt.Print(degisken_1, degisken_2, degisken_3)

}
