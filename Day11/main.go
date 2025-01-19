package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type state []int;

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input file path as first argument")
		return
	}
	inputFilePath := os.Args[1]
	input, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	for i := 0; i < 25; i++ {
		input = tick(input)
	}
	fmt.Println(len(input))
}

func tick(stones state) state {
	newStones := make(state, 0, len(stones))
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if hasEvenDigits(stone) {
			firstHalf, secondHalf := splitEvenDigits(stone)
			newStones = append(newStones, firstHalf, secondHalf)
		} else {
			newStones = append(newStones, stone * 2024)
		}
	}
	return newStones
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}

func splitEvenDigits(n int) (int, int) {
	digits := countDigits(n)

	divisor := 1
	for i := 0; i < digits/2; i++ {
		divisor *= 10
	}

	firstHalf := n / divisor
	secondHalf := n % divisor

	return firstHalf, secondHalf
}

func hasEvenDigits(n int) bool {
	if n == 0 {
		return false
	}
	return countDigits(n)%2 == 0
}

func readInput(inputFilePath string) (state, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stones state

	if scanner.Scan() {
		line := scanner.Text()
		var num int
		for _, field := range strings.Fields(line) {
			fmt.Sscanf(field, "%d", &num)
			stones = append(stones, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stones, nil
}
