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
		newPoints := resonancePoints(antennas)
		newPoints = removeInvalidPoints(newPoints, inputs.X, inputs.Y)
		// plotPoints(newPoints, inputs.X, inputs.Y)
		points = append(points, newPoints...)
	}
	points = removeDuplicatePoints(points)
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

func resonancePoints(antennas []point) []point {
	var result []point
	// Get all pairwise combinations
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			// Get anti-nodes for this pair and add them to result
			antiNodePoints := antiNodes(antennas[i], antennas[j])
			result = append(result, antiNodePoints...)
		}
	}

	return result
}

func antiNodes(a1, a2 point) []point {
	midPoint := point{X: (a1.X + a2.X) / 2, Y: (a1.Y + a2.Y) / 2}
	v1 := point{X: a1.X - midPoint.X, Y: a1.Y - midPoint.Y}
	v2 := point{X: a2.X - midPoint.X, Y: a2.Y - midPoint.Y}
	p1 := point{X: a1.X + v1.X*2, Y: a1.Y + v1.Y*2}
	p2 := point{X: a2.X + v2.X*2, Y: a2.Y + v2.Y*2}
	antiNodes := make([]point, 0)
	if p1 != a1 && p1 != a2 {
		antiNodes = append(antiNodes, p1)
	}
	if p2 != a1 && p2 != a2 {
		antiNodes = append(antiNodes, p2)
	}
	return antiNodes
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
		// Check bounds and ensure coordinates are whole numbers
		if p.X >= 0 && p.X < float64(maxX) && p.Y >= 0 && p.Y < float64(maxY) &&
			float64(int(p.X)) == p.X && float64(int(p.Y)) == p.Y {
			validPoints = append(validPoints, p)
		}
	}
	return validPoints
}
