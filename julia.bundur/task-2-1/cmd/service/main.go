package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	scanner.Split(bufio.ScanWords)
	next := func() string {
		scanner.Scan()
		return scanner.Text()
	}
	depts, _ := strconv.Atoi(next())
	for i := 0; i < depts; i++ {
		staff, _ := strconv.Atoi(next())
		minTemp := 15
		maxTemp := 30
		for j := 0; j < staff; j++ {
			sign := next()
			var numText string
			if sign == ">=" || sign == "<=" {
				numText = next()
			} else {
				numText = sign[2:]
				sign = sign[:2]
			}
			temp, _ := strconv.Atoi(strings.TrimSpace(numText))
			if sign == ">=" {
				if temp > minTemp {
					minTemp = temp
				}
			} else {
				if temp < maxTemp {
					maxTemp = temp
				}
			}
			if minTemp > maxTemp {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, minTemp)
			}
		}
	}
}
