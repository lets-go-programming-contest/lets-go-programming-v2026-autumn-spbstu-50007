package main

import "fmt"

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil || departments < 1 || departments > 1000 {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departments {
		if !processDepartment() {
			return
		}
	}
}

func processDepartment() bool {
	var employees int

	_, err := fmt.Scan(&employees)
	if err != nil || employees < 1 || employees > 1000 {
		fmt.Println("Invalid number of employees")

		return false
	}

	lowest := 15
	highest := 30

	for range employees {
		var (
			sign        string
			temperature int
		)

		_, err = fmt.Scan(&sign, &temperature)
		if err != nil {
			fmt.Println("Invalid temperature requirement")

			return false
		}

		switch sign {
		case ">=":
			if temperature > lowest {
				lowest = temperature
			}
		case "<=":
			if temperature < highest {
				highest = temperature
			}
		default:
			fmt.Println("Invalid operator")

			return false
		}

		if lowest > highest {
			fmt.Println(-1)
		} else {
			fmt.Println(lowest)
		}
	}

	return true
}
