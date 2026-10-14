package main

import (
	"fmt"
)

func alogosNeMnogoIf() int {
	var empl int
	low := 15
	high := 30
	_, err := fmt.Scanln(&empl)
	if err != nil || empl < 1 || empl > 1000 {
		fmt.Println("Invalid number of employees")
		return 1
	}
	for range empl {
		var operat string
		var temp int
		_, err = fmt.Scanln(&operat, &temp)
		if err != nil {
			fmt.Println("Invalid temperature requirement")
			return 1
		}
		if operat == ">=" {
			low = max(low, temp)
		} else if operat == "<=" {
			high = min(high, temp)
		} else {
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
	_, err := fmt.Scanln(&deprt)
	if err != nil || deprt < 1 || deprt > 1000 {
		fmt.Println("Invalid number of departments")

		return
	}
	for range deprt {
		if alogosNeMnogoIf() != 0 {
			return
		}
	}

}

