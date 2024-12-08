package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	inputFilePath := "input2.txt"
	data, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	lines := allHorizontalLines(data)
	lines = append(lines, allVerticalLines(data)...)
	lines = append(lines, allDiagonalLines(data)...)
	forward := regexp.MustCompile(`XMAS`)
	backward := regexp.MustCompile(`SAMX`)
	sum := 0
	for _, line := range lines {
		sum += countOverlappingMatches(forward, line)
		sum += countOverlappingMatches(backward, line)
	}
	fmt.Println(sum)
}

type state struct {
	x_count, m_count, a_count, s_count                     int
	last_x_index, last_m_index, last_a_index, last_s_index int
}

func initMemo(s string) []state {
	memo := make([]state, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		var prevMemo state
		if i == len(s)-1 {
			prevMemo = state{}
		} else {
			prevMemo = memo[i+1]
		}
		newSate := state{x_count: prevMemo.x_count, m_count: prevMemo.m_count, a_count: prevMemo.a_count, s_count: prevMemo.s_count, last_x_index: prevMemo.last_x_index, last_m_index: prevMemo.last_m_index, last_a_index: prevMemo.last_a_index, last_s_index: prevMemo.last_s_index}
		switch s[i] {
		case 'X':
			newSate.x_count++
			newSate.last_x_index = i
		case 'M':
			newSate.m_count++
			newSate.last_m_index = i
		case 'A':
			newSate.a_count++
			newSate.last_a_index = i
		case 'S':
			newSate.s_count++
			newSate.last_s_index = i
		}
		memo[i] = newSate
	}
	return memo
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
