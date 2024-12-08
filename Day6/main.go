package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x, y int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type State struct {
	obstructions []Pos
	guardPos Pos
	direction Direction
	visited map[Pos]bool
	bounds Pos
}

func main() {
	inputFilePath := "input2.txt"
	state, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	nextState, shouldContinue := tick(state)
	for shouldContinue {
		nextState, shouldContinue = tick(nextState)
	}
	sum := 0
	for _, visited := range nextState.visited {
		if visited {
			sum++
		}
	}
	fmt.Println(sum)
}

func tick(currentState State) (State, bool) {
	nextPos := nextStep(currentState.guardPos, currentState.direction)
	nextDirection := currentState.direction
	if !posInBounds(nextPos, currentState) {
		return currentState, false
	}
	if isObstructed(nextPos, currentState) {
		nextDirection = (currentState.direction + 1) % 4
		nextPos = nextStep(currentState.guardPos, nextDirection)
		if !posInBounds(nextPos, currentState) {
			return currentState, false
		}
	}
	visited := currentState.visited
	visited[nextPos] = true
	nextState := State{
		obstructions: currentState.obstructions,
		guardPos: nextPos,
		direction: nextDirection,
		visited: visited,
		bounds: currentState.bounds,
	}
	return nextState, true
}

func isObstructed(pos Pos, state State) bool {
	for _, obstruction := range state.obstructions {
		if pos == obstruction {
			return true
		}
	}
	return false
}

func posInBounds(pos Pos, state State) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.x < state.bounds.x && pos.y < state.bounds.y
}

func nextStep(currentPos Pos, direction Direction) Pos {
	switch direction {
		case North:
			return Pos{currentPos.x, currentPos.y - 1}
		case East:
			return Pos{currentPos.x + 1, currentPos.y}
		case South:
			return Pos{currentPos.x, currentPos.y + 1}
		case West:
			return Pos{currentPos.x - 1, currentPos.y}
		default:
			panic(fmt.Sprintf("Invalid direction: %d", direction))
	}
}


func readInput(inputFilePath string) (State, error) {
	file, err := os.Open(inputFilePath)
	if (err != nil) {
		return State{}, err
	}
	defer file.Close()

	var data State
	scanner := bufio.NewScanner(file)
	y := 0
	XMax := 0
	for scanner.Scan() {
		line := scanner.Text()
        if line == "" {
            continue
        }
		XMax = len(line)
		for x, char := range line {
			currentPos := Pos{x, y}
			switch char {
				case '^':
					data.guardPos = currentPos
					data.direction = North
				case '>':
					data.guardPos = currentPos
					data.direction = East
				case 'v':
					data.guardPos = currentPos
					data.direction = South
				case '<':
					data.guardPos = currentPos
					data.direction = West
				case '#':
					data.obstructions = append(data.obstructions, currentPos)
				case '.':
					// Do nothing
				default:
					return State{}, fmt.Errorf("invalid character '%c' at position (%d, %d)", char, x, y)
			}
		}
		y++
	}

	data.bounds = Pos{XMax, y}

	data.visited = make(map[Pos]bool)
	data.visited[data.guardPos] = true
	if err := scanner.Err(); err != nil {
		return State{}, err
	}

	return data, nil
}
