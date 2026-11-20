package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	first, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	scanner.Scan()
	second, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	scanner.Scan()
	op := scanner.Text()
	switch op {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		if second == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(first / second)
	default:
		fmt.Println("Invalid operation")
	}
}
