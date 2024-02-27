package main

import "fmt"

func dataTypes() {
	var first uint8 = 200
	var second int8 = -10
	var decimal float32 = 12.5
	var condition bool = true
	var message = `Salam kenal : ),
	Selamat datang di "Hacktiv8",
	Mari belajar "Scalable Web Service With Go"`

	fmt.Printf(`Tipe data variable first : %T`, first)
	fmt.Printf(`Tipe data variable second : %T`, second)
	fmt.Printf(`Tipe data variable decimal : %T`, decimal)
	fmt.Printf("decimal : %.2f", decimal)
	fmt.Printf("decimal : %f", decimal)
	fmt.Printf(`Tipe data variable condition : %T`, condition)
	fmt.Printf("Is it true? : %t", condition)
	fmt.Print(message)

}
