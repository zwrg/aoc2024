package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func day3part1() int {
	return 0
}

func day3part2() int {
	return 0
}

func main() {
	filePath := filepath.Join("data", "2024", "day3.txt")

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

	for scanner.Scan() {
		line := scanner.Text()
	}

	println(day3part1())
	println(day3part2())
}
