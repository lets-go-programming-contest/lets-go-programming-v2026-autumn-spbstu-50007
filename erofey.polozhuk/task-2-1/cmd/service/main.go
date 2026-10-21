package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidEmployees = errors.New("invalid number of employees")
	ErrInvalidOperator  = errors.New("invalid operator")
)

func departmentProcess() error {
	var empl int

	low := 15
	high := 30

	if _, err := fmt.Scan(&empl); err != nil {
		return fmt.Errorf("reading employees count: %w", err)
	}

	if empl < 1 || empl > 1000 {
		return fmt.Errorf("%w: %d", ErrInvalidEmployees, empl)
	}

	for range empl {
		var (
			operat string
			temp   int
		)

		if _, err := fmt.Scan(&operat, &temp); err != nil {
			return fmt.Errorf("invalid temperature requirement: %w", err)
		}

		switch operat {
		case ">=":
			low = max(low, temp)
		case "<=":
			high = min(high, temp)
		default:
			return fmt.Errorf("%w: %s", ErrInvalidOperator, operat)
		}

		if low > high {
			fmt.Println(-1)
		} else {
			fmt.Println(low)
		}
	}

	return nil
}

func main() {
	var deprt int

	if _, err := fmt.Scan(&deprt); err != nil {
		fmt.Println("reading departments count:", err)

		return
	}

	if deprt < 1 || deprt > 1000 {
		fmt.Println("invalid number of departments")

		return
	}

	for range deprt {
		if err := departmentProcess(); err != nil {
			fmt.Println(err)

			return
		}
	}
}
