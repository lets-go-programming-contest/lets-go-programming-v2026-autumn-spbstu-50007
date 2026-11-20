package main

import "fmt"

func main() {
	var number1, number2 int
	var operator string

	if _, err := fmt.Scanln(&number1); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scanln(&number2); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err := fmt.Scanln(&operator); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operator {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		if number2 == 0 {
			fmt.Println("Division by zero")
			return
		} else {
			fmt.Println(number1 / number2)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
