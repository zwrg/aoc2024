package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Data struct {
	target  int
	numbers []int
}

func solveEquation(target int, eqs []int, concat bool) int {
	stack := [][]int{eqs}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(current) == 1 {
			if current[0] == target {
				return target
			}
			continue
		}
		newNums := current[2:]

		stack = append(stack, append([]int{current[0] + current[1]}, newNums...))
		stack = append(stack, append([]int{current[0] * current[1]}, newNums...))
		if concat {
			n, _ := strconv.Atoi(strconv.Itoa(current[0]) + strconv.Itoa(current[1]))
			stack = append(stack, append([]int{n}, newNums...))
		}
	}

	return 0
}

func day7part1(equations []Data) int {
	sum := 0
	for _, equation := range equations {
		sum += solveEquation(equation.target, equation.numbers, false)
	}

	return sum
}

func day7part2(equations []Data) int {
	sum := 0
	for _, equation := range equations {
		sum += solveEquation(equation.target, equation.numbers, true)
	}

	return sum
}

func main() {
	filePath := filepath.Join("data", "2024", "day7.txt")

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var equations []Data

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ": ")
		target, _ := strconv.Atoi(split[0])

		numberString := strings.Split(split[1], " ")
		numbers := make([]int, len(numberString))

		for i, num := range numberString {
			numbers[i], _ = strconv.Atoi(num)
		}

		equations = append(equations, Data{target, numbers})
	}

	println(day7part1(equations))
	println(day7part2(equations))
}
