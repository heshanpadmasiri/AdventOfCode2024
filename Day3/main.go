package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFilePath := "input2.txt"
	input, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		num1, err1 := strconv.Atoi(match[1])
		num2, err2 := strconv.Atoi(match[2])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to integer:", err1, err2)
			continue
		}
		sum += (num1 * num2)
	}
	fmt.Println(sum)
}


func readInput(inputFilePath string) (string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := ""
	for scanner.Scan() {
		line := scanner.Text()
        if line == "" {
            continue
        }
		result += line
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return result, nil
}
