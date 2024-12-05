package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func day5part1(data [][]int, compareFn func(a, b int) int) int {
	sum := 0
	for _, v := range data {
		if slices.IsSortedFunc(v, compareFn) {
			sum += v[len(v)/2]
		}
	}
	return sum
}

func day5part2(data [][]int, compareFn func(a, b int) int) int {
	sum := 0
	for _, v := range data {
		if !slices.IsSortedFunc(v, compareFn) {
			slices.SortFunc(v, compareFn)
			sum += v[len(v)/2]
		}
	}
	return sum
}

func main() {
	filePath := filepath.Join("data", "2024", "day5.txt")

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

	input, _ := os.ReadFile(filePath)

	split := strings.Split(string(input), "\n\n")
	_rules, _data := strings.Split(split[0], "\n"), strings.Split(split[1], "\n")

	rules := make([][]int, len(_rules))
	data := make([][]int, len(_data))

	for i, rule := range _rules {
		r := strings.Split(rule, "|")
		num, _ := strconv.Atoi(r[0])
		numRule, _ := strconv.Atoi(r[1])
		rules[i] = []int{num, numRule}
	}

	for i, operation := range _data {
		for _, op := range strings.Split(operation, ",") {
			var n, _ = strconv.Atoi(op)
			data[i] = append(data[i], n)
		}
	}

	compareFunction := func(a, b int) int {
		for _, r := range rules {
			if r[0] == a && r[1] == b {
				return -1
			}
		}
		return 0
	}

	println(day5part1(data, compareFunction))
	println(day5part2(data, compareFunction))
}
