package main

import "fmt"

func conditionInit() {
	var currentYear = 2024

	if age := currentYear - 2000; age > 20 {
		fmt.Println("Age is above 20")
	} else {
		fmt.Println("Age is below 20")
	}

}
