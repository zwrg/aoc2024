package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func day1part1(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	var diffSum = 0
	for i := range left {
		diffSum += int(math.Abs(float64(left[i] - right[i])))
	}
	return diffSum
}

func day1part2(left []int, right []int) int {
	countMap := make(map[int]int)
	for _, number := range right {
		countMap[number]++
	}
	sum := 0
	for _, number := range left {
		sum += number * countMap[number]
	}
	return sum
}

func main() {
	filePath := filepath.Join("data", "2024", "day1.txt")

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
	var leftColumn []int
	var rightColumn []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting numbers:", err1, err2)
			continue
		}

		leftColumn = append(leftColumn, num1)
		rightColumn = append(rightColumn, num2)
	}

	println(day1part1(leftColumn, rightColumn))
	println(day1part2(leftColumn, rightColumn))
}
