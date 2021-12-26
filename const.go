package main 

import "fmt"

const sabit_1 = "sabit 1"


// birden fazla const olduğunda aşağıdaki gibi süslü parantezle yaoabiliriz
const {

	sabit_2 = "sabit 2"
	sabit_3 = "sabit 3"

}


// constlar sıralı şekilde artan değerlerse  iota kullanılıyor ve değerler birer birer artıyor
const {

	sabit_4 = iota
	sabit_5
	sabit_6


}



func main() {

	fmt.Print(sabit_1, sabit_2, sabit_3, sabit_4, sabit_5, sabit_6)


}