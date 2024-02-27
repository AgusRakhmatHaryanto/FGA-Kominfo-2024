package main

import "fmt"

func switchInit() {
	var score = 8

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
	}

	var score2 = 6

	switch {
	case score2 == 8:
		fmt.Println("Perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("Not bad")
		fallthrough
	case score2 < 5:
		fmt.Println("Need to study more")
	default:
		fmt.Println("study harder")
		fmt.Println("you need to study more")
	}
}
