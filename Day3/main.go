package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type insnKind int

const (
	mul insnKind = iota
	do
	dont
)

type insn struct {
	kind insnKind
	index int
	val int
}

func main() {
	inputFilePath := "input2.txt"
	input, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	insns := getMulInsn(input)
	insns = append(insns, getDoInsn(input)...)
	insns = append(insns, getDontInsn(input)...)
	sort.Slice(insns, func(i, j int) bool {
		return insns[i].index < insns[j].index
	})
	enabled := true
	sum := 0
	for _, insn := range insns {
		if insn.kind == mul && enabled {
			sum += insn.val
		} else if insn.kind == dont {
			enabled = false
		} else if insn.kind == do {
			enabled = true
		}
	}
	fmt.Println(sum)
}

func getDontInsn(input string) []insn {
	reRegex := regexp.MustCompile(`don't\(\)`)
	matches := reRegex.FindAllStringIndex(input, -1)
	insns := make([]insn, 0)
	for _, match := range matches {
		insns = append(insns, insn{kind: dont, index: match[0], val: 0})
	}
	return insns
}

func getDoInsn(input string) []insn {
	reRegex := regexp.MustCompile(`do\(\)`)
	matches := reRegex.FindAllStringIndex(input, -1)
	insns := make([]insn, 0)
	for _, match := range matches {
		insns = append(insns, insn{kind: do, index: match[0], val: 0})
	}
	return insns
}

func getMulInsn(input string) []insn {

	reRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := reRegex.FindAllStringSubmatchIndex(input, -1)
	insns := make([]insn, 0)
	for _, match := range matches {
		num1, err1 := strconv.Atoi(input[match[2]:match[3]])
		num2, err2 := strconv.Atoi(input[match[4]:match[5]])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to integer:", err1, err2)
			continue
		}
		insns = append(insns, insn{kind: mul, index: match[0], val: num1 * num2})
	}
	return insns
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
