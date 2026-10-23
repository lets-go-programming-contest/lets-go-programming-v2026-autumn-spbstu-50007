package main

import (
	"fmt"
	"os"
)

func checkResult(minTemp, maxTemp int) {
	if minTemp > maxTemp {
		fmt.Println(-1)
	} else {
		fmt.Println(minTemp)
	}
}

func processDepartment() {
	var employees int
	if _, err := fmt.Scan(&employees); err != nil {
		os.Exit(1)
	}

	minTemp := 15

	maxTemp := 30

	for range employees {
		var operator string

		var temp int

		if _, err := fmt.Scan(&operator, &temp); err != nil {
			os.Exit(1)
		}

		switch operator {
		case ">=":
			if temp > minTemp {
				minTemp = temp
			}
		case "<=":
			if temp < maxTemp {
				maxTemp = temp
			}
		case ">":
			if temp+1 > minTemp {
				minTemp = temp + 1
			}
		case "<":
			if temp-1 < maxTemp {
				maxTemp = temp - 1
			}
		}

		checkResult(minTemp, maxTemp)
	}
}

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		return
	}

	for range departments {
		processDepartment()
	}
}
