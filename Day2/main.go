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
			printReport(report)
			sum++
		}
	}
	fmt.Println(sum)
}

func printReport(report Report) {
	for i, val := range report.vals {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}

func isSafe(report Report) bool {
	vals := report.vals
	if isSafeInner(vals) {
		return true;
	}
	for i := 0; i < len(report.vals); i++ {
		newVals := skipIndex(vals, i)
		if isSafeInner(newVals) {
			isSafeInner(newVals)
			return true
		}
	}
	return false;
}

func skipIndex(vals []int, index int) []int {
	newVals := make([]int, 0)
	for i, val := range vals {
		if i != index {
			newVals = append(newVals, val)
		}
	}
	return newVals
}

func isIncreasing(vals []int) bool {
	if len(vals) < 2 {
		return true
	}
	for i := 0; i < len(vals) - 1; i++ {
		if vals[i] >= vals[i + 1] || vals[i + 1] - vals[i] > 3 {
			return false
		}
	}
	return true
}

func isDecreasing(vals []int) bool {
	if len(vals) < 2 {
		return true
	}
	for i := 0; i < len(vals) - 1; i++ {
		if vals[i] <= vals[i + 1] || vals[i] - vals[i + 1] > 3 {
			return false
		}
	}
	return true
}

func isMonotonic(vals []int) bool {
	if len(vals) == 0 {
		return false
	}
	if len(vals) == 1 {
		return true
	}
	diff := vals[0] - vals[1]
	if diff == 0 {
		return false
	}
	for i := 2; i < len(vals); i++ {
		newDiff := vals[i-1] - vals[i]
		if newDiff == 0 {
			return false
		}
		if diff > 0 && newDiff < 0 || diff < 0 && newDiff > 0 {
			return false
		}
	}
	return true
}

func isSafeDiff(vals []int) bool {
	for i := 0; i < len(vals) - 1; i++ {
		diff := vals[i] - vals[i+1]
		if abs(diff) > 3 {
			return false
		}
	}
	return true
}

func isSafeInner(vals []int) bool {
	return isIncreasing(vals) || isDecreasing(vals);
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
