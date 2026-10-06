package main

import (
	"fmt"
)

func main() {
	var (
		a, b, x int
		operand string
	)

	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Invalid first operator", err)
		return
	}
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Invalid second operator", err)
		return
	}
	_, err = fmt.Scanln(&operand)
	if err != nil {
		fmt.Println("Invalid operation", err)
		return
	}

	switch operand {
	case "+":
		x = a + b
		fmt.Println(x)
	case "-":
		x = a - b
		fmt.Println(x)
	case "*":
		x = a * b
		fmt.Println(x)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		} else {
			x = a / b
			fmt.Println(x)
		}

	default:
		fmt.Println("Invalid operation")
	}
}
