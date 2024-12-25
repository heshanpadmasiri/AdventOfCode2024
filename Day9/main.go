package main

import (
	"fmt"
	"os"
)

type block int

const empty block = -1

type blockMemo struct {
	ptr int
	len int
}

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
	blocks, freeSpaces, files := parseInput(input)
	for i := len(files) - 1; i >= 0; i-- {
		compact2(blocks, i, freeSpaces, files)
	}
	// printState(blocks)
	fmt.Println(checkSum(blocks))
}

func compact2(blocks []block, fileIndex int, freeSpaces, files []blockMemo) {
	targetIndex := -1
	for i := 0; i < len(freeSpaces); i++ {
		if freeSpaces[i].ptr > files[fileIndex].ptr {
			return
		}
		if freeSpaces[i].len >= files[fileIndex].len {
			targetIndex = i
			break
		}
	}
	if targetIndex == -1 {
		return
	}
	size := files[fileIndex].len
	freeBlockPtr := freeSpaces[targetIndex].ptr
	targetPtr := files[fileIndex].ptr
	for i := 0; i < size; i++ {
		blocks[freeBlockPtr+i] = blocks[targetPtr+i]
		blocks[targetPtr+i] = empty
	}
	freeSpaces[targetIndex].len -= size
	freeSpaces[targetIndex].ptr += size
}

func validateState(blocks []block, emptyPtr, blockPtr int) {
	if blocks[emptyPtr] != empty {
		panic("emptyPtr is not empty")
	}
	if blocks[blockPtr] == empty {
		panic("blockPtr is empty")
	}
	for i := 0; i < emptyPtr; i++ {
		if blocks[i] == empty {
			panic("block before emptyPtr is empty")
		}
	}
	for i := blockPtr + 1; i < len(blocks); i++ {
		if blocks[i] != empty {
			panic("block after blockPtr is not empty")
		}
	}
}

func checkSum(blocks []block) int {
	sum := 0
	for i, b := range blocks {
		if b != empty {
			sum += int(b) * (i)
		}
	}
	return sum
}

func getNextEmptyPtr(blocks []block, emptyPtr int) int {
	for i := emptyPtr; i < len(blocks); i++ {
		if blocks[i] == empty {
			return i
		}
	}
	return -1
}

func getNextBlockPtr(blocks []block, blockPtr int) int {
	for i := blockPtr; i >= 0; i-- {
		if blocks[i] != empty {
			return i
		}
	}
	return -1
}

func compact(blocks []block, emptyPtr, blockPtr int) (bool, int, int) {
	if emptyPtr > blockPtr {
		return false, -1, -1
	}
	if blocks[blockPtr] == empty {
		panic("blockPtr is empty")
	}
	if blocks[emptyPtr] != empty {
		panic("emptyPtr is not empty")
	}
	blocks[emptyPtr] = blocks[blockPtr]
	blocks[blockPtr] = empty
	if blocks[emptyPtr] == empty {
		panic("emptyPtr is not filled")
	}
	newEmptyPtr := getNextEmptyPtr(blocks, emptyPtr+1)
	newBlockPtr := getNextBlockPtr(blocks, blockPtr-1)
	return true, newEmptyPtr, newBlockPtr
}

func parseInput(input []uint8) ([]block, []blockMemo, []blockMemo) {
	blocks := make([]block, 0)
	freeSpaces := make([]blockMemo, 0)
	files := make([]blockMemo, 0)
	blockId := 0
	isBlock := true
	index := 0
	for _, v := range input {
		if isBlock {
			for i := 0; i < int(v); i++ {
				blocks = append(blocks, block(blockId))
			}
			files = append(files, blockMemo{ptr: index, len: int(v)})
			blockId++
		} else {
			for i := 0; i < int(v); i++ {
				blocks = append(blocks, empty)
			}
			freeSpaces = append(freeSpaces, blockMemo{ptr: index, len: int(v)})
		}
		isBlock = !isBlock
		index += int(v)
	}
	return blocks, freeSpaces, files
}

func printState(blocks []block) {
	for _, b := range blocks {
		if b == empty {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", b)
		}
	}
	fmt.Println()
}

func readInput(inputFilePath string) ([]uint8, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the entire file content
	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		return nil, err
	}

	// Convert string of digits to uint8 array
	// Trim any whitespace or newlines
	cleanContent := []byte(string(content))
	var result []uint8

	for _, ch := range cleanContent {
		// Skip newlines or other whitespace
		if ch >= '0' && ch <= '9' {
			// Convert ASCII digit to actual number
			result = append(result, ch-'0')
		}
	}

	return result, nil
}
