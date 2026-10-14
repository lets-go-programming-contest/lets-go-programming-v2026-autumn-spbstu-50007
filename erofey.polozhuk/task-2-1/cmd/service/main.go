package main

import (
	"fmt"
)

func alogosNeMnogoIf() int {
	var empl int

	low := 15
	high := 30

	if _, err := fmt.Scan(&empl); err != nil || empl < 1 || empl > 1000 {
		fmt.Println("Invalid number of employees")

		return 1
	}

	for range empl {
		var (
			operat string
			temp   int
		)

		if _, err := fmt.Scan(&operat, &temp); err != nil {
			fmt.Println("Invalid temperature requirement")

			return 1
		}

		switch operat {
		case ">=":
			low = max(low, temp)
		case "<=":
			high = min(high, temp)
		default:
			fmt.Println("Invalid operator")

			return 1
		}

		if low > high {
			fmt.Println(-1)
		} else {
			fmt.Println(low)
		}
	}

	return 0
}

func main() {
	var deprt int

	if _, err := fmt.Scan(&deprt); err != nil || deprt < 1 || deprt > 1000 {
		fmt.Println("Invalid number of departments")

		return
	}

	for range deprt {
		if alogosNeMnogoIf() != 0 {
			return
		}
	}
}
