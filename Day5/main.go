package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	ordering map[int][]int
	updates [][]int
}

func main() {
	inputFilePath := "input2.txt"
	data, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	sum := 0
	for _, update := range data.updates {
		if isValid(data.ordering, update) {
			middleValue := update[len(update) / 2]
			sum += middleValue
		}
	}
	fmt.Println(sum)
}

func isValid(ordering map[int][]int, update []int) bool {
	seen := make(map[int]bool)
	for _, num := range update {
		after := ordering[num]
		for _, a := range after {
			if seen[a] {
				return false
			}
		}
		seen[num] = true
	}
	return true
}

func readInput(inputFilePath string) (Data, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return Data{}, err
	}
	defer file.Close()

	data := Data{
		ordering: make(map[int][]int),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
        if line == "" {
            break
        }
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return Data{}, fmt.Errorf("invalid line format: %s", line)
		}
		lh, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return Data{}, fmt.Errorf("error converting left part: %w", err)
		}
		rh, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return Data{}, fmt.Errorf("error converting right part: %w", err)
		}
		data.ordering[lh] = append(data.ordering[lh], rh)
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		sNum := strings.Split(line, ",")
		nums := make([]int, 0)
		for _, num := range sNum {
			n, err := strconv.Atoi(num)
			if err != nil {
				return Data{}, fmt.Errorf("error converting number: %w", err)
			}
			nums = append(nums, n)
		}	
		data.updates = append(data.updates, nums)

	}

	if err := scanner.Err(); err != nil {
		return Data{}, err
	}

	return data, nil
}
