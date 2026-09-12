package main

import (
	"fmt"
)

func main() {
	var (
		a        int
		b        int
		operator string
		result   int
	)

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&operator)
	if err != nil || (operator != "+" && operator != "-" && operator != "*" && operator != "/") {
		fmt.Println("Invalid operation")
		return
	}

	if operator == "/" && b == 0 {
		fmt.Println("Division by zero")
		return
	}

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	fmt.Println(result)
}
