package main

import "fmt"

// 3 boyutlu integer dizisi olusturmak için:
var arr_1 [3]int

// içerisine deger atamak için

var arr_2 = [5]int{1, 2, 3, 4, 5}

func main() {

	// func içinde array tanımı yapmak için make kullanılıyor
	arr_3 := make([]int, 3)

	arr_3[1] = 2

	// arr_1 e değer atanmadığı için 000 bastı
	// bu int tipler için0 string için boş string olarak döner
	fmt.Println(arr_1, arr_2, arr_3)

	// arrayin boyutunu veren fonk
	fmt.Printf("arr_1 len:%d \n", len(arr_1))

}
