package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	inputFilePath := "input2.txt"
	lines, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	sum := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == 'M' || lines[y][x] == 'S' {
				if isMatch(lines, x, y) {
					// for i := 0; i < 3; i++ {
					// 	for j := 0; j < 3; j++ {
					// 		fmt.Print(string(lines[y+i][x+j]))
					// 	}
					// 	fmt.Println()
					// }
					// fmt.Println()
					sum++
				}
			}
		}
	}
	fmt.Println(sum)
}

func isMatch(lines[] string, x, y int) bool {
	width := len(lines[0])
	if x + 2 >= width || y + 2 >= len(lines) {
		return false
	}
	if lines[y+1][x+1] != 'A' {
		return false
	}
	var e1 byte;
	if lines[y][x] == 'M' {
		e1 = 'S'
	} else {
		e1 = 'M'
	}
	if lines[y+2][x+2] != e1 {
		return false
	}

	var e2 byte;
	if lines[y][x+2] == 'M' {
		e2 = 'S'
	} else if lines[y][x+2] == 'S' {
		e2 = 'M'
	} else {
		return false
	}
	if lines[y+2][x] != e2 {
		return false
	}
	return true
}

func countOverlappingMatches(re *regexp.Regexp, s string) int {
	found := re.FindAllStringSubmatch(s, -1)
	return len(found)
}

func allHorizontalLines(lines []string) []string {
	result := make([]string, len(lines))
	copy(result, lines)
	return result
}

func allVerticalLines(lines []string) []string {
	if len(lines) == 0 {
		return []string{}
	}

	columnCount := len(lines[0])
	verticalLines := make([]string, columnCount)

	for _, line := range lines {
		for i, char := range line {
			verticalLines[i] += string(char)
		}
	}

	return verticalLines
}

func allDiagonalLines(lines []string) []string {
	if len(lines) == 0 {
		return []string{}
	}

	rows := len(lines)
	cols := len(lines[0])
	diagonals := make([]string, 0)

	// Top-left to bottom-right diagonals
	for startCol := 0; startCol < cols; startCol++ {
		diagonal := ""
		for i := 0; i < rows && startCol+i < cols; i++ {
			diagonal += string(lines[i][startCol+i])
		}
		if len(diagonal) > 0 {
			diagonals = append(diagonals, diagonal)
		}
	}

	for startRow := 1; startRow < rows; startRow++ {
		diagonal := ""
		for i := 0; i < cols && startRow+i < rows; i++ {
			diagonal += string(lines[startRow+i][i])
		}
		if len(diagonal) > 0 {
			diagonals = append(diagonals, diagonal)
		}
	}

	// Top-right to bottom-left diagonals
	for startCol := 0; startCol < cols; startCol++ {
		diagonal := ""
		for i := 0; i < rows && startCol-i >= 0; i++ {
			diagonal += string(lines[i][startCol-i])
		}
		if len(diagonal) > 0 {
			diagonals = append(diagonals, diagonal)
		}
	}

	for startRow := 1; startRow < rows; startRow++ {
		diagonal := ""
		for i := 0; i < cols && startRow+i < rows; i++ {
			diagonal += string(lines[startRow+i][cols-1-i])
		}
		if len(diagonal) > 0 {
			diagonals = append(diagonals, diagonal)
		}
	}

	return diagonals
}

func readInput(inputFilePath string) ([]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
