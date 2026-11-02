package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	bufferSize = 1024 * 1024
	minLimit   = 15
	maxLimit   = 30
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	defer func() {
		_ = writer.Flush()
	}()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, bufferSize), bufferSize)
	scanner.Split(bufio.ScanWords)

	next := func() string {
		scanner.Scan()

		return scanner.Text()
	}

	depts, _ := strconv.Atoi(next())

	for range depts {
		staff, _ := strconv.Atoi(next())
		minTemp := minLimit
		maxTemp := maxLimit

		for range staff {
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
				minTemp = max(minTemp, temp)
			} else {
				maxTemp = min(maxTemp, temp)
			}

			if minTemp > maxTemp {
				_, _ = fmt.Fprintln(writer, -1)
			} else {
				_, _ = fmt.Fprintln(writer, minTemp)
			}
		}
	}
}
