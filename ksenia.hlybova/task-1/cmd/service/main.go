package main

import (
	"fmt"
	"strconv"
)

func main() {
	var aStr, bStr, op string

	_, err := fmt.Scanln(&aStr)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&bStr)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&op)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	a, err := strconv.Atoi(aStr)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = a / b
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)
}
