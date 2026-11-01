package main

import "fmt"

var errOp = fmt.Errorf("operator must be %q or %q", ">=", "<=")

func getTemperature(oper string, temp, minT, maxT int) (int, int, int, error) {
	switch oper {
	case ">=":
		if temp > minT {
			minT = temp
		}
	case "<=":
		if temp < maxT {
			maxT = temp
		}
	default:
		return minT, maxT, 0, fmt.Errorf("%w; received %q", errOp, oper)
	}

	if minT > maxT {
		return minT, maxT, -1, nil
	}

	return minT, maxT, minT, nil
}

func main() {
	var deps int

	_, err := fmt.Scan(&deps)
	if err != nil {
		fmt.Printf("reading departments count: %s\n", err.Error())

		return
	}

	if deps < 1 || deps > 1000 {
		fmt.Println("Invalid number of departments")

		return
	}

	for range deps {
		var emps int

		_, err = fmt.Scan(&emps)
		if err != nil {
			fmt.Printf("reading employees count: %s\n", err.Error())

			return
		}

		if emps < 1 || emps > 1000 {
			fmt.Println("Invalid number of employees")

			return
		}

		minT := 15
		maxT := 30

		for range emps {
			var oper string

			var temp int

			_, err = fmt.Scan(&oper, &temp)
			if err != nil {
				fmt.Printf("reading temperature requirement: %s\n", err.Error())

				return
			}

			newMin, newMax, res, err := getTemperature(oper, temp, minT, maxT)
			if err != nil {
				fmt.Printf("getting temperature: %s\n", err.Error())

				return
			}

			minT = newMin
			maxT = newMax

			fmt.Println(res)
		}
	}
}
