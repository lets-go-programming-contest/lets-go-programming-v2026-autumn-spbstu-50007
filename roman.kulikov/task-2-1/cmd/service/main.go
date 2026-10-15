package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)

		low := minTemp
		high := maxTemp

		for j := 0; j < k; j++ {
			var sign string
			var vl int
			fmt.Scan(&sign, &vl)

			switch sign {
			case ">=":
				if vl > low {
					low = vl
				}
			case "<=":
				if vl < high {
					high = vl
				}
			}

			if low <= high {
				fmt.Println(low)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
