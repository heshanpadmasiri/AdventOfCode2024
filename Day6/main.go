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
	obstructions map[Pos]bool
	guardPos Pos
	direction Direction
	visited map[Pos][]Direction
	bounds Pos
}

func main() {
	inputFilePath := "input2.txt"
	state, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	sum := 0
	for y := 0; y < state.bounds.y; y++ {
		for x := 0; x < state.bounds.x; x++ {
			pos := Pos{x, y}
			if pos != state.guardPos && !isObstructed(pos, state) {
				newState := addObstructions(state, pos)
				if isLoop(newState) {
					sum++
				}
			}
		}
	}
	fmt.Println(sum)
}

type Cnt int

const (
	Continue Cnt = iota
	Stop
	Loop
)

func addObstructions(state State, pos Pos) State {
	newObstructions := make(map[Pos]bool)
	for k, v := range state.obstructions {
		newObstructions[k] = v
	}
	newObstructions[pos] = true

	newVisited := make(map[Pos][]Direction)
	for k, v := range state.visited {
		directionsCopy := make([]Direction, len(v))
		copy(directionsCopy, v)
		newVisited[k] = directionsCopy
	}
	return State{
		obstructions: newObstructions,
		guardPos: state.guardPos,
		direction: state.direction,
		visited: newVisited,
		bounds: state.bounds,
	}
}

func isLoop(state State) bool {
	nextState, cnt := tick(state)
	for cnt == Continue {
		nextState, cnt = tick(nextState)
	}
	return cnt == Loop
}

func tick(currentState State) (State, Cnt) {
	nextPos := nextStep(currentState.guardPos, currentState.direction)
	nextDirection := currentState.direction
	if !posInBounds(nextPos, currentState) {
		return currentState, Stop
	}
	count := 0
	for isObstructed(nextPos, currentState) {
		count++
		if count > 4 {
			fmt.Print("Error: guard is stuck at ", currentState.guardPos, " with direction ", currentState.direction, "\n")	
		}
		nextDirection = (nextDirection + 1) % 4
		nextPos = nextStep(currentState.guardPos, nextDirection)
		if !posInBounds(nextPos, currentState) {
			return currentState, Stop
		}
	}
	visited := currentState.visited
	for _, direction := range visited[nextPos] {
		if direction == nextDirection {
			return currentState, Loop
		}
	}
	visited[nextPos] = append(visited[nextPos], nextDirection)
	nextState := State{
		obstructions: currentState.obstructions,
		guardPos: nextPos,
		direction: nextDirection,
		visited: visited,
		bounds: currentState.bounds,
	}
	return nextState, Continue
}

func isObstructed(pos Pos, state State) bool {
	return state.obstructions[pos]
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
	data.obstructions = make(map[Pos]bool)
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
					data.obstructions[currentPos] = true
				case '.':
					// Do nothing
				default:
					return State{}, fmt.Errorf("invalid character '%c' at position (%d, %d)", char, x, y)
			}
		}
		y++
	}

	data.bounds = Pos{XMax, y}

	data.visited = make(map[Pos][]Direction)
	data.visited[data.guardPos] = []Direction{data.direction}
	if err := scanner.Err(); err != nil {
		return State{}, err
	}

	return data, nil
}
