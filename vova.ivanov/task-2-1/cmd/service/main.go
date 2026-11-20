package main

import "fmt"

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Printf("reading departments count: %s\n", err.Error())

		return
	}

	if departments < 1 || departments > 1000 {
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
	if err != nil {
		fmt.Printf("reading employees count: %s\n", err.Error())

		return false
	}

	if employees < 1 || employees > 1000 {
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
			fmt.Printf("reading temperature requirement: %s\n", err.Error())

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
