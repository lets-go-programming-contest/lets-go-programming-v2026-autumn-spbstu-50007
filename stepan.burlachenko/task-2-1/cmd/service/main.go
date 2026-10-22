package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func entryData(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return -1, -1, errors.New("failed to read line")
	}
	line := scanner.Text()

	nTemp, kTemp, success := strings.Cut(line, " ")
	if !success {
		return -1, -1, errors.New("failed to cut line")
	}

	n, err := strconv.Atoi(nTemp)
	if err != nil {
		return -1, -1, err
	}

	k, err := strconv.Atoi(kTemp)
	if err != nil {
		return -1, -1, err
	}

	return n, k, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, k, err := entryData(scanner)
	if err != nil {
		fmt.Println("-1")
		return
	}

	for i := 0; i < n; i++ {
		minTemp := 15
		maxTemp := 30

		for j := 0; j < k; j++ {
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

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
