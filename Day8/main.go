package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	X float64
	Y float64
}

type input struct {
	X        int
	Y        int
	antennas map[string][]point
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
	points := make([]point, 0)
	for _, antennas := range inputs.antennas {
		// fmt.Println(antenna)
		newPoints := resonancePoints(antennas, inputs.X, inputs.Y)
		newPoints = removeInvalidPoints(newPoints, inputs.X, inputs.Y)
		points = append(points, newPoints...)
	}
	points = removeDuplicatePoints(points)
	// plotPoints(points, inputs.X, inputs.Y)
	fmt.Println(len(points))
}

func plotPoints(points []point, maxX, maxY int) {
	grid := make([][]string, maxY)
	for i := range grid {
		grid[i] = make([]string, maxX)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	for _, p := range points {
		grid[int(p.Y)][int(p.X)] = "#"
	}
	for _, row := range grid {
		fmt.Println(strings.Join(row, ""))
	}
}

func removeDuplicatePoints(points []point) []point {
	seen := make(map[point]bool)
	var uniquePoints []point
	for _, p := range points {
		if !seen[p] {
			seen[p] = true
			uniquePoints = append(uniquePoints, p)
		}
	}
	return uniquePoints
}

func resonancePoints(antennas []point, X, Y int) []point {
	var result []point
	// Get all pairwise combinations
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			// Get anti-nodes for this pair and add them to result
			antiNodePoints := antiNodes(antennas[i], antennas[j], X, Y)
			result = append(result, antiNodePoints...)
		}
	}

	return result
}

func antiNodes(a1, a2 point, X, Y int) []point {
	y_1 := a1.Y
	y_2 := a2.Y
	x_1 := a1.X
	x_2 := a2.X
	var result []point
	for x := 0; x < X; x++ {
		x_bar := float64(x)
		a_1 := y_1 * (x_bar - x_2)
		a_2 := y_2 * (x_bar - x_1)
		a := a_1 - a_2
		b := x_1 - x_2
		y_bar := a / b
		point := point{X: x_bar, Y: y_bar}
		if isPointValid(point, X, Y) {
			result = append(result, point)
		}
	}
	return result
}

func readInput(inputFilePath string) (input, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return input{}, err
	}
	defer file.Close()

	// Read all lines to get dimensions and content
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return input{}, err
	}

	// Create input structure
	result := input{
		Y:        len(lines),
		X:        len(lines[0]),
		antennas: make(map[string][]point),
	}

	// Scan through the grid to find antennas
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				antenna := string(char)
				result.antennas[antenna] = append(result.antennas[antenna], point{X: float64(x), Y: float64(y)})
			}
		}
	}

	return result, nil
}

func removeInvalidPoints(points []point, maxX, maxY int) []point {
	var validPoints []point
	for _, p := range points {
		if isPointValid(p, maxX, maxY) {
			validPoints = append(validPoints, p)
		}
	}
	return validPoints
}

func isPointValid(p point, maxX, maxY int) bool {
	return p.X >= 0 && p.X < float64(maxX) && p.Y >= 0 && p.Y < float64(maxY) &&
		float64(int(p.X)) == p.X && float64(int(p.Y)) == p.Y
}
