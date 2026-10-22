package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readConstraint(scanner *bufio.Scanner) (string, int, bool) {
	if !scanner.Scan() {
		return "", 0, false
	}
	token := scanner.Text()

	var operator string
	var valStr string

	if strings.HasPrefix(token, ">=") {
		operator = ">="
		valStr = strings.TrimPrefix(token, ">=")
	} else if strings.HasPrefix(token, "<=") {
		operator = "<="
		valStr = strings.TrimPrefix(token, "<=")
	} else {
		operator = token
		valStr = ""
	}

	if valStr == "" {
		if !scanner.Scan() {
			return "", 0, false
		}
		valStr = scanner.Text()
	}

	value, err := strconv.Atoi(valStr)
	if err != nil {
		return "", 0, false
	}

	return operator, value, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}

	numDepartments, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	for i := 0; i < numDepartments; i++ {
		if !scanner.Scan() {
			break
		}
		numEmployees, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}

		minTemp := 15
		maxTemp := 30

		for j := 0; j < numEmployees; j++ {
			operator, value, ok := readConstraint(scanner)
			if !ok {
				break
			}

			if operator == ">=" {
				if value > minTemp {
					minTemp = value
				}
			} else if operator == "<=" {
				if value < maxTemp {
					maxTemp = value
				}
			}

			if minTemp > maxTemp {
				fmt.Println("-1")
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
