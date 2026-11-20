package main

import "fmt"

func main() {
	var fir int
	var sec int
	var operation string

	_, err := fmt.Scanf("%d\n", &fir)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanf("%d\n", &sec)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(fir + sec)
	case "-":
		fmt.Println(fir - sec)
	case "*":
		fmt.Println(fir * sec)
	case "/":
		if sec == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(fir / sec)
	default:
		fmt.Println("Invalid operation")
	}
}
