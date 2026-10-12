package main

import "fmt"

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil || departments < 1 || departments > 1000 {
		fmt.Println("Invalid number of departments")
		return
	}

	for i := 0; i < departments; i++ {
		var workers int

		_, err = fmt.Scan(&workers)
		if err != nil || workers < 1 || workers > 1000 {
			fmt.Println("Invalid number of employees")
			return
		}

		lowest := 15
		highest := 30

		for j := 0; j < workers; j++ {
			var sign string
			var temperature int

			_, err = fmt.Scan(&sign, &temperature)
			if err != nil {
				fmt.Println("Invalid temperature requirement")
				return
			}

			if sign == ">=" {
				if temperature > lowest {
					lowest = temperature
				}
			} else if sign == "<=" {
				if temperature < highest {
					highest = temperature
				}
			} else {
				fmt.Println("Invalid operator")
				return
			}

			if lowest > highest {
				fmt.Println(-1)
			} else {
				fmt.Println(lowest)
			}
		}
	}
}
