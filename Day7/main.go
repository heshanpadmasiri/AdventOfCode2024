package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type input struct {
	target int
	values []int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input file path as first argument")
		return
	}
	inputFilePath := os.Args[1]
	inputs, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	sum := 0
	for _, input := range inputs {
		if isValid(input) {
			sum += input.target
		}
	}
	fmt.Println(sum)
}

func isValid(input input) bool {
	return isValidInner(input.target, input.values[0], 0, input.values[1:], add) ||
		isValidInner(input.target, input.values[0], 1, input.values[1:], mul) ||
		isValidInner(input.target, input.values[0], 0, input.values[1:], concat)
}

func isValidInner(target, current, currentSum int, remainingValues []int, operation operation) bool {
	value := currentSum
	if operation == add {
		value += current
	} else if operation == mul {
		value *= current
	} else if operation == concat {
		currentStr := fmt.Sprint(current)
		currentSumStr := fmt.Sprint(currentSum)
		value, _ = strconv.Atoi(currentSumStr + currentStr)
		// var digits int
		// if current != 0 {
		// 	digits = int(math.Floor(math.Log10(float64(current)))) + 1
		// } else {
		// 	digits = 1
		// }
		// value = currentSum*int(math.Pow10(digits)) + current
	}
	if value > target {
		return false
	}
	if len(remainingValues) == 0 {
		return value == target
	}
	return isValidInner(target, remainingValues[0], value, remainingValues[1:], add) ||
		isValidInner(target, remainingValues[0], value, remainingValues[1:], mul) ||
		isValidInner(target, remainingValues[0], value, remainingValues[1:], concat)
}

type operation int

const (
	add operation = iota
	mul
	concat
)

func readInput(inputFilePath string) ([]input, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return []input{}, err
	}
	defer file.Close()

	var inputs []input
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var target int
		var values []int

		// Parse the line using fmt.Sscanf for the target
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		if _, err := fmt.Sscanf(parts[0], "%d", &target); err != nil {
			return nil, fmt.Errorf("error parsing target: %v", err)
		}

		// Parse the values
		valueStrs := strings.Fields(parts[1])
		for _, str := range valueStrs {
			var val int
			if _, err := fmt.Sscanf(str, "%d", &val); err != nil {
				return nil, fmt.Errorf("error parsing value: %v", err)
			}
			values = append(values, val)
		}

		inputs = append(inputs, input{
			target: target,
			values: values,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return inputs, nil
}
