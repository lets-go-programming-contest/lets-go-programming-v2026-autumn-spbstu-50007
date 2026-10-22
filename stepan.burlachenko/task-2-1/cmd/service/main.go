package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	ErrFailedToReadLine = errors.New("failed to read token")
)

func entryData(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return -1, -1, ErrFailedToReadLine
	}
	nTemp := scanner.Text()

	departmentsCount, err := strconv.Atoi(nTemp)
	if err != nil {
		return -1, -1, fmt.Errorf("failed to parse departments count: %w", err)
	}

	if !scanner.Scan() {
		return -1, -1, ErrFailedToReadLine
	}
	kTemp := scanner.Text()

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
		operator := scanner.Text()

		if !scanner.Scan() {
			break
		}
		temp, err := strconv.Atoi(scanner.Text())
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
			fmt.Println("-1")
		} else {
			fmt.Println(minTemp)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

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
