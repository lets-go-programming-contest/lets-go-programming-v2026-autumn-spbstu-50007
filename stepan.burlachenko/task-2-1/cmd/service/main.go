package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}

	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	if !scanner.Scan() {
		return
	}

	k, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	for range n {
		minT := 15
		maxT := 30

		for range k {
			if !scanner.Scan() {
				break
			}

			op := scanner.Text()
			valStr := ""

			if strings.HasPrefix(op, ">=") {
				valStr = op[2:]
				op = ">="
			} else if strings.HasPrefix(op, "<=") {
				valStr = op[2:]
				op = "<="
			}

			if valStr == "" {
				if !scanner.Scan() {
					break
				}

				valStr = scanner.Text()
			}

			val, err := strconv.Atoi(valStr)
			if err != nil {
				fmt.Println("-1")

				continue
			}

			if op == ">=" {
				if val > minT {
					minT = val
				}
			} else if op == "<=" {
				if val < maxT {
					maxT = val
				}
			}

			if minT > maxT {
				fmt.Println("-1")
			} else {
				fmt.Println(minT)
			}
		}
	}
}
