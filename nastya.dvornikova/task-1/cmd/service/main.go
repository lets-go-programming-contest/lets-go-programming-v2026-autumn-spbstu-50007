package main

import (
	"fmt"
)

func main() {
	var x, y int
	var op string

	if _, err := fmt.Scan(&x); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err := fmt.Scan(&y); err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, err := fmt.Scan(&op); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	if op == "/" && y == 0 {
		fmt.Println("Division by zero")
		return
	}

	var res int
	switch op {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "*":
		res = x * y
	case "/":
		res = x / y
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(res)
}
