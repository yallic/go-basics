package main 

import "fmt"

const sabit_1 = "sabit 1"


// birden fazla const oldugunda asagıdaki gibi suslu parantezle yaoabiliriz
const {

	sabit_2 = "sabit 2"
	sabit_3 = "sabit 3"

}


// constlar sıralı sekilde artan degerlerse  iota kullanılıyor ve degerler birer birer artıyor
const {

	sabit_4 = iota
	sabit_5
	sabit_6


}



func main() {

	fmt.Print(sabit_1, sabit_2, sabit_3, sabit_4, sabit_5, sabit_6)


}