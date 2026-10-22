package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readOperand(input *bufio.Scanner) (int, bool) {
	if !input.Scan() {
		return 0, false
	}

	value, err := strconv.Atoi(strings.TrimSpace(input.Text()))
	return value, err == nil
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	first, ok := readOperand(input)
	if !ok {
		fmt.Println("Invalid first operand")
		return
	}

	second, ok := readOperand(input)
	if !ok {
		fmt.Println("Invalid second operand")
		return
	}

	if !input.Scan() {
		fmt.Println("Invalid operation")
		return
	}

	switch strings.TrimSpace(input.Text()) {
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
