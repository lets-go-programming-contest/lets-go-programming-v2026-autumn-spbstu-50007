package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processDepartment(scanner *bufio.Scanner, numEmployees int) {
	minTemp := 15
	maxTemp := 30

	for range numEmployees {
		operator, value, ok := readConstraint(scanner)

		if !ok {
			break
		}

		switch operator {
		case ">=":
			if value > minTemp {
				minTemp = value
			}
		case "<=":
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

func readConstraint(scanner *bufio.Scanner) (string, int, bool) {
	if !scanner.Scan() {
		return "", 0, false
	}

	token := scanner.Text()

	var operator string
	var valStr string

	switch {
	case strings.HasPrefix(token, ">="):
		operator = ">="
		valStr = strings.TrimPrefix(token, ">=")
	case strings.HasPrefix(token, "<="):
		operator = "<="
		valStr = strings.TrimPrefix(token, "<=")
	default:
		operator = token
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

	for range numDepartments {
		if !scanner.Scan() {
			break
		}

		numEmployees, err := strconv.Atoi(scanner.Text())

		if err != nil {
			break
		}

		processDepartment(scanner, numEmployees)
	}
}
