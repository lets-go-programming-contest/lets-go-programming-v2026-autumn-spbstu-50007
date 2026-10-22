package main

import "fmt"

func main() {
	var argument1, argument2 int
	var operat string

	if _, err := fmt.Scan(&argument1); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err := fmt.Scan(&argument2); err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, err := fmt.Scan(&operat); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operat {
	case "+":
		fmt.Println(argument1 + argument2)
	case "-":
		fmt.Println(argument1 - argument2)
	case "*":
		fmt.Println(argument1 * argument2)
	case "/":
		if argument2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(argument1 / argument2)
	default:
		fmt.Println("Invalid operation")
	}
}
