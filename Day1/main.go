package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	lhs []int
	rhs []int
}

func main() {
	inputFilePath := "input2.txt"
	data, err := readInput(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
    // sort.Ints(data.lhs)
    // sort.Ints(data.rhs)

    // sum := 0
    // for i := 0; i < len(data.lhs); i++ {
    //     sum +=  abs(data.lhs[i] - data.rhs[i])
    // }
    // fmt.Println(sum)
    memo := make(map[int]int)
    for i := 0; i < len(data.rhs); i++ {
        val := data.rhs[i]
		memo[val]++
    }
    sum := 0;
    for i := 0; i < len(data.lhs); i++ {
        sum += data.lhs[i] * memo[data.lhs[i]]
    }
    fmt.Println(sum)
}

func abs(x int) int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}

func readInput(inputFilePath string) (Data, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return Data{}, err
	}
	defer file.Close()

	var data Data
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
        if line == "" {
            continue
        }
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			lhs, _ := strconv.Atoi(numbers[0])
			rhs, _ := strconv.Atoi(numbers[1])
			data.lhs = append(data.lhs, lhs)
			data.rhs = append(data.rhs, rhs)
		} else {
            return Data{}, fmt.Errorf("invalid input: %s", line)
        }
	}

	if err := scanner.Err(); err != nil {
		return Data{}, err
	}

	return data, nil
}
