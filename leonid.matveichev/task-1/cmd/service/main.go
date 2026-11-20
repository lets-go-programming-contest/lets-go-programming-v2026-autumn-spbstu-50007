package main

import "fmt"

func main() {
	var firstNumber, secondNumber int
	var mathematicalOperation string

	if _, err := fmt.Scan(&firstNumber); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err := fmt.Scan(&secondNumber); err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, err := fmt.Scan(&mathematicalOperation); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch mathematicalOperation {
	case "+":
		fmt.Println(firstNumber + secondNumber)
	case "-":
		fmt.Println(firstNumber - secondNumber)
	case "*":
		fmt.Println(firstNumber * secondNumber)
	case "/":
		if secondNumber == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(firstNumber / secondNumber)
	default:
		fmt.Println("Invalid operation")
	}
}
