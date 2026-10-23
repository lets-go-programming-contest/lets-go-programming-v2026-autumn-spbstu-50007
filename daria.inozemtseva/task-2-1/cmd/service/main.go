package main

import "fmt"

func checkDepartment(empCount int) {
	minTemp := 15
	maxTemp := 30

	for range empCount {
		var symbol string

		var temp int

		if _, err := fmt.Scan(&symbol, &temp); err != nil {
			return
		}

		if symbol == ">=" {
			if temp > minTemp {
				minTemp = temp
			}
		} else if symbol == "<=" {
			if temp < maxTemp {
				maxTemp = temp
			}
		}

		if minTemp <= maxTemp {
			fmt.Println(minTemp)
		} else {
			fmt.Println("-1")
		}
	}
}

func main() {
	var deptCount int

	if _, err := fmt.Scan(&deptCount); err != nil {
		return
	}

	if deptCount < 1 || deptCount > 1000 {
		return
	}

	for range deptCount {
		var empCount int

		if _, err := fmt.Scan(&empCount); err != nil {
			return
		}

		if empCount < 1 || empCount > 1000 {
			return
		}

		checkDepartment(empCount)
	}
}
