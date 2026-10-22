package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrFailedToReadLine = errors.New("failed to read line")
	ErrFailedToCutLine  = errors.New("failed to cut line")
)

func entryData(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return -1, -1, ErrFailedToReadLine
	}

	line := scanner.Text()

	nTemp, kTemp, success := strings.Cut(line, " ")
	if !success {
		return -1, -1, ErrFailedToCutLine
	}

	departmentsCount, err := strconv.Atoi(nTemp)
	if err != nil {
		return -1, -1, fmt.Errorf("failed to parse departments count: %w", err)
	}

	employeesPerDepartment, err := strconv.Atoi(kTemp)
	if err != nil {
		return -1, -1, fmt.Errorf("failed to parse employees count: %w", err)
	}

	return departmentsCount, employeesPerDepartment, nil
}

func processDepartment(scanner *bufio.Scanner, employeesPerDepartment int) {
	minTemp := 15
	maxTemp := 30

	for range employeesPerDepartment {
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		operator, numStr, success := strings.Cut(line, " ")
		if !success {
			fmt.Println("-1")

			continue
		}

		temp, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("-1")

			continue
		}

		switch operator {
		case ">=":
			if temp > minTemp {
				minTemp = temp
			}
		case "<=":
			if temp < maxTemp {
				maxTemp = temp
			}
		default:
			fmt.Println("-1")

			continue
		}

		if minTemp > maxTemp {
			fmt.Print("-1 ")
		} else {
			fmt.Print(minTemp, " ")
		}
	}

	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	departmentsCount, employeesPerDepartment, err := entryData(scanner)
	if err != nil {
		fmt.Println("-1")

		return
	}

	for range departmentsCount {
		processDepartment(scanner, employeesPerDepartment)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
