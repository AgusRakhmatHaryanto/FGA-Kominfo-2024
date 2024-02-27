package main

import "fmt"

func operatorsInit() {
	const a = 10
	fmt.Print(a, "\n")

	// a := 10

	// fmt.Print(a)
	//error

	var value = 2 * (3 - 1)

	fmt.Print(value)

	var firstCondition bool = 2 > 3
	var secondCondition bool = "Hello" == "hello"
	var thirdCondition bool = 10 != 11
	var fourthCondition bool = 10 <= 10

	fmt.Print(firstCondition, secondCondition, thirdCondition, fourthCondition)

	var right = true
	var wrong = false

	var wrongAndRight = right && wrong
	fmt.Print("\n",wrongAndRight)

	var wrongOrRight = right || wrong
	fmt.Print(wrongOrRight)

	var wrongReverse = !wrong
	fmt.Print(wrongReverse)
}
