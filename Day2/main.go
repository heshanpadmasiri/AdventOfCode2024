package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	reports []Report
}

type Report struct {
	vals []int
}

func main() {
	inputFilePath := "input2.txt"
	input, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	sum := 0
	for _, report := range input.reports {
		if isSafe(report) {
			sum++
		}
	}
	fmt.Println(sum)
}

func isSafe(report Report) bool {
	if len(report.vals) == 0 {
		return true
	}
	var increasing bool
	var diff = report.vals[0] - report.vals[1]
	if diff == 0 || abs(diff) > 3 {
		return false
	}
	if diff > 0 {
		increasing = false
	} else {
		increasing = true
	}
	for i := 2; i < len(report.vals); i++ {
		diff = report.vals[i-1] - report.vals[i]
		if diff == 0 || abs(diff) > 3 {
			return false
		}
		if increasing && diff > 0 {
			return false
		}
		if !increasing && diff < 0 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func readInput(inputFilePath string) (Input, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return Input{}, err
	}
	defer file.Close()

	var input Input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
        if line == "" {
            continue
        }
		numbers := strings.Fields(line)
		var report Report
		if len(numbers) > 0 {
			for _, num := range numbers {
				val, err := strconv.Atoi(num)
				if err != nil {
					return Input{}, err
				}
				report.vals = append(report.vals, val)
			}
		} else {
            return Input{}, fmt.Errorf("invalid input: %s", line)
        }
		input.reports = append(input.reports, report)
	}
	if err := scanner.Err(); err != nil {
		return Input{}, err
	}

	return input, nil
}
