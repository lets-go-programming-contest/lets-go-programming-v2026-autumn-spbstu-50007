package main

import (
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	}
	return 0, fmt.Errorf("Division by zero")
}
func Summa(a, b int) int {
	return a + b
}
func Diff(a, b int) int {
	return a - b
}
func Multip(a, b int) int {
	return a * b
}
func Operation(a, b int, opertion string) (int, error) {
	switch opertion {
	case "+":
		res := Summa(a, b)
		return res, nil
	case "-":
		res := Diff(a, b)
		return res, nil
	case "*":
		res := Multip(a, b)
		return res, nil
	case "/":
		res, err := Divide(a, b)
		return res, err
	default:
		return 0, fmt.Errorf("Invalid operation")
	}
}

func main() {
	var (
		a         int
		b         int
		operation string
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
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}
	result, err := Operation(a, b, operation)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
