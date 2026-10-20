package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	minTemperature = 15
	maxTemperature = 30
	maxCount       = 1000
)

var (
	errInvalidCount    = errors.New("count must be between 1 and 1000")
	errInvalidOperator = errors.New("operator must be <= or >=")
)

func readCount(reader *bufio.Reader, field string) (int, error) {
	var count int

	if _, err := fmt.Fscan(reader, &count); err != nil {
		return 0, fmt.Errorf("reading %s count: %w", field, err)
	}

	if count < 1 || count > maxCount {
		return 0, fmt.Errorf("invalid %s count: %w", field, errInvalidCount)
	}

	return count, nil
}

func processDepartment(reader *bufio.Reader, writer *bufio.Writer) error {
	employees, err := readCount(reader, "employees")
	if err != nil {
		return err
	}

	lower := minTemperature
	upper := maxTemperature

	for range employees {
		var (
			operator    string
			temperature int
		)

		if _, err := fmt.Fscan(reader, &operator, &temperature); err != nil {
			return fmt.Errorf("reading temperature requirement: %w", err)
		}

		switch operator {
		case ">=":
			if temperature > lower {
				lower = temperature
			}
		case "<=":
			if temperature < upper {
				upper = temperature
			}
		default:
			return fmt.Errorf("%w: %q", errInvalidOperator, operator)
		}

		answer := lower
		if lower > upper {
			answer = -1
		}

		if _, err := fmt.Fprintln(writer, answer); err != nil {
			return fmt.Errorf("writing temperature: %w", err)
		}
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	departments, err := readCount(reader, "departments")
	if err != nil {
		fmt.Println(err)

		return
	}

	for range departments {
		err = processDepartment(reader, writer)
		if err != nil {
			break
		}
	}

	if flushErr := writer.Flush(); flushErr != nil {
		fmt.Printf("writing temperatures: %s\n", flushErr.Error())
	}

	if err != nil {
		fmt.Println(err)
	}
}
