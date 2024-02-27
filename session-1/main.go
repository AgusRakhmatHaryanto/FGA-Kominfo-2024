package main

import (
	"fmt"
)

func main() {

	//varible with data type
	var name string = "Joko"
	var age int = 30
	email := "joko@koko"
	var buah1, buah2, buah3 string = "Apel", "Jeruk", "Nanas"
	var harga1, harga2, harga3 int
	harga1, harga2, harga3 = 1000, 2000, 3000
	var campus string
	var jalan, nomor, kota = "Jalan Raya", 1, "Jakarta"
	campus = "Universitas Nusa Mandiri"

	//underscore variable
	var kodePos string
	_= kodePos


	fmt.Println("Nama saya :", name)
	fmt.Println("Umur saya :", age)
	fmt.Println("Email saya :", email)
	fmt.Println("Saya tinggal di :", jalan, nomor, kota)
	fmt.Println("Saya berasal dari :", campus)
	fmt.Printf("Buah 1 : %s, Buah 2 : %s, Buah 3 : %s\n", buah1, buah2, buah3)
	fmt.Printf("Harga 1 : %d, Harga 2 : %d, Harga 3 : %d\n", harga1, harga2, harga3)
	fmt.Printf(`Hello, nama saya %s saya berumyur %d tahun saya berasal dari %s`, name, age, campus)

	//variable without data type
	// var name = "Joko"
	// var age = 30

	// fmt.Println("Nama saya :", name)
	// fmt.Println("Umur saya :", age)

	fmt.Println("Hello, World!")
	initDB()
}
