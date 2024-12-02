package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func day2IsLineSafe(line []int) bool {
	isIncreasing, isDecreasing := true, true
	for i := 0; i < len(line)-1; i++ {
		diff := line[i] - line[i+1]
		if diff > 0 {
			isIncreasing = false
		}
		if diff < 0 {
			isDecreasing = false
		}
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}
	}
	return isIncreasing || isDecreasing
}

func day2part1(numbers [][]int) int {
	count := 0
	for _, line := range numbers {
		if day2IsLineSafe(line) {
			count++
		}
	}
	return count
}

func day2part2(numbers [][]int) int {
	count := 0
	for _, line := range numbers {
		if day2IsLineSafe(line) {
			count++
			continue
		}
		for i := 0; i < len(line); i++ {
			modifiedLine := append(append([]int{}, line[:i]...), line[i+1:]...)
			if day2IsLineSafe(modifiedLine) {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	filePath := filepath.Join("data", "2024", "day2.txt")

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

	var numbers [][]int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var intParts []int

		for _, str := range parts {
			num, _ := strconv.Atoi(str)
			intParts = append(intParts, num)
		}
		numbers = append(numbers, intParts)
	}

	println(day2part1(numbers))
	println(day2part2(numbers))
}
