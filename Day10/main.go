package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	X int
	Y int
}

type height uint8

type grid [][]height

type memo map[point][]point

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
	sum := 0
	memo := make(memo)
	for _, start := range startPositions(input) {
		sum += numberOfTrails(input, memo, start)
	}
	fmt.Println(sum)
}

func numberOfTrails(grid grid, memo memo, currentPosition point) int {
	reachable := reachableEnds(grid, memo, currentPosition)
	// Create a map to track unique points
	uniquePoints := make(map[point]bool)
	for _, p := range reachable {
		uniquePoints[p] = true
	}
	return len(uniquePoints)
}

func reachableEnds(grid grid, memo memo, currentPosition point) []point {
	currentHeight := grid[currentPosition.Y][currentPosition.X]
	if currentHeight == 9 {
		return []point{currentPosition}
	}
	if val, exists := memo[currentPosition]; exists {
		return val
	}
	var reachable []point
	for _, move := range possibleMoves(grid, currentPosition) {
		reachable = append(reachable, reachableEnds(grid, memo, move)...)
	}
	memo[currentPosition] = reachable
	return reachable
}

func possibleMoves(grid grid, currentPosition point) []point {
	var moves []point
	currentHeight := grid[currentPosition.Y][currentPosition.X]
	
	// Define possible directions: up, down, left, right
	directions := []point{
		{X: 0, Y: -1}, // up
		{X: 0, Y: 1},  // down
		{X: -1, Y: 0}, // left
		{X: 1, Y: 0},  // right
	}
	
	// Check each direction
	for _, dir := range directions {
		newPos := point{
			X: currentPosition.X + dir.X,
			Y: currentPosition.Y + dir.Y,
		}
		
		// Check if new position is within grid bounds
		if newPos.Y < 0 || newPos.Y >= len(grid) || 
		   newPos.X < 0 || newPos.X >= len(grid[0]) {
			continue
		}
		
		// Check if new height is exactly one more than current height
		newHeight := grid[newPos.Y][newPos.X]
		if newHeight == currentHeight+1 {
			moves = append(moves, newPos)
		}
	}
	
	return moves
}

func startPositions(grid grid) []point {
	var starts []point

	// Iterate through each row and column
	for y := range grid {
		for x := range grid[y] {
			// If height is 0, add the position to starts
			if grid[y][x] == 0 {
				starts = append(starts, point{X: x, Y: y})
			}
		}
	}

	return starts
}

func readInput(inputFilePath string) (grid, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result grid
	scanner := bufio.NewScanner(file)

	// Read file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue // Skip empty lines
		}

		// Convert each character in the line to a height
		row := make([]height, len(line))
		for i, ch := range line {
			if ch < '0' || ch > '9' {
				return nil, fmt.Errorf("invalid character in input: %c", ch)
			}
			row[i] = height(ch - '0')
		}
		result = append(result, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
