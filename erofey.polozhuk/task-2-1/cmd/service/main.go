package main

import (
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n < 1 || n > 1000 {
		fmt.Println("Invalid number of departments")
		return
	}
	for range n {
		var k int
		_, err := fmt.Scan(&k)
		if err != nil || k < 1 || k > 1000 {
			fmt.Println("Invalid number of employees")
			return
		}
		var temp int
		var operator string
		le, ri := 15, 30
		for range k {
			_, err := fmt.Scan(&operator, &temp)
			if err != nil {
				fmt.Println("Invalid temperature requirement")
				return
			}
			switch operator {
			case ">=":
				le = max(le, temp)
			case "<=":
				ri = min(ri, temp)
			}
			if le <= ri {
				fmt.Println(le)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
