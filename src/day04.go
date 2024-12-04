package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func day4checkForXMAS(lines [][]string, i int, j int, marks [][]bool) int {
	count := 0
	if i+3 < len(lines) && lines[i][j] == "X" && lines[i+1][j] == "M" && lines[i+2][j] == "A" && lines[i+3][j] == "S" {
		marks[i][j], marks[i+1][j], marks[i+2][j], marks[i+3][j] = true, true, true, true
		count++
	}
	if i-3 >= 0 && lines[i][j] == "X" && lines[i-1][j] == "M" && lines[i-2][j] == "A" && lines[i-3][j] == "S" {
		marks[i][j], marks[i-1][j], marks[i-2][j], marks[i-3][j] = true, true, true, true
		count++
	}
	if j+3 < len(lines[i]) && lines[i][j] == "X" && lines[i][j+1] == "M" && lines[i][j+2] == "A" && lines[i][j+3] == "S" {
		marks[i][j], marks[i][j+1], marks[i][j+2], marks[i][j+3] = true, true, true, true
		count++
	}
	if j-3 >= 0 && lines[i][j] == "X" && lines[i][j-1] == "M" && lines[i][j-2] == "A" && lines[i][j-3] == "S" {
		marks[i][j], marks[i][j-1], marks[i][j-2], marks[i][j-3] = true, true, true, true
		count++
	}

	if i+3 < len(lines) && j+3 < len(lines[i]) && lines[i][j] == "X" && lines[i+1][j+1] == "M" && lines[i+2][j+2] == "A" && lines[i+3][j+3] == "S" {
		marks[i][j], marks[i+1][j+1], marks[i+2][j+2], marks[i+3][j+3] = true, true, true, true
		count++
	}
	if i+3 < len(lines) && j-3 >= 0 && lines[i][j] == "X" && lines[i+1][j-1] == "M" && lines[i+2][j-2] == "A" && lines[i+3][j-3] == "S" {
		marks[i][j], marks[i+1][j-1], marks[i+2][j-2], marks[i+3][j-3] = true, true, true, true
		count++
	}
	if i-3 >= 0 && j+3 < len(lines[i]) && lines[i][j] == "X" && lines[i-1][j+1] == "M" && lines[i-2][j+2] == "A" && lines[i-3][j+3] == "S" {
		marks[i][j], marks[i-1][j+1], marks[i-2][j+2], marks[i-3][j+3] = true, true, true, true
		count++
	}
	if i-3 >= 0 && j-3 >= 0 && lines[i][j] == "X" && lines[i-1][j-1] == "M" && lines[i-2][j-2] == "A" && lines[i-3][j-3] == "S" {
		marks[i][j], marks[i-1][j-1], marks[i-2][j-2], marks[i-3][j-3] = true, true, true, true
		count++
	}

	return count
}

func day4checkForMAS(lines [][]string, i int, j int, marks [][]bool) int {
	count := 0
	if i+1 >= len(lines) || j+1 >= len(lines[i]) || i-1 < 0 || j-1 < 0 {
		return 0
	}

	if lines[i-1][j-1] == "M" && lines[i+1][j+1] == "S" && lines[i+1][j-1] == "M" && lines[i-1][j+1] == "S" ||
		lines[i-1][j+1] == "M" && lines[i+1][j-1] == "S" && lines[i+1][j+1] == "M" && lines[i-1][j-1] == "S" {
		// horizontal
		marks[i][j], marks[i-1][j+1], marks[i+1][j-1], marks[i-1][j-1], marks[i+1][j+1] = true, true, true, true, true
		count++
	}
	if lines[i-1][j-1] == "M" && lines[i+1][j+1] == "S" && lines[i-1][j+1] == "M" && lines[i+1][j-1] == "S" ||
		lines[i+1][j-1] == "M" && lines[i+-1][j+1] == "S" && lines[i+1][j+1] == "M" && lines[i-1][j-1] == "S" {
		// vertical
		marks[i][j], marks[i-1][j+1], marks[i+1][j-1], marks[i-1][j-1], marks[i+1][j+1] = true, true, true, true, true
		count++
	}

	return count
}

func day4part1(lines [][]string) int {
	marks := make([][]bool, len(lines))
	for i := range marks {
		marks[i] = make([]bool, len(lines[i]))
		for j := range marks[i] {
			marks[i][j] = false
		}
	}

	count := 0
	for i, _ := range lines {
		for j, _ := range lines[i] {
			if lines[i][j] == "X" {
				count += day4checkForXMAS(lines, i, j, marks)
			}
		}
	}

	//for i, _ := range lines {
	//	for j, _ := range lines[i] {
	//		if marks[i][j] {
	//			fmt.Print(lines[i][j], " ")
	//		} else {
	//			fmt.Print(". ")
	//		}
	//	}
	//	fmt.Println()
	//}

	return count
}

func day4part2(lines [][]string) int {
	marks := make([][]bool, len(lines))
	for i := range marks {
		marks[i] = make([]bool, len(lines[i]))
		for j := range marks[i] {
			marks[i][j] = false
		}
	}

	count := 0
	for i, _ := range lines {
		for j, _ := range lines[i] {
			if lines[i][j] == "A" {
				count += day4checkForMAS(lines, i, j, marks)
			}
		}
	}

	//for i, _ := range lines {
	//	for j, _ := range lines[i] {
	//		if marks[i][j] {
	//			fmt.Print(lines[i][j], " ")
	//		} else {
	//			fmt.Print(". ")
	//		}
	//	}
	//	fmt.Println()
	//}

	return count
}

func main() {
	filePath := filepath.Join("data", "2024", "day4.txt")

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

	var lines [][]string

	var i = 0
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for j, char := range line {
			fmt.Print(string(char), "(", i, ",", j, ") ")
			row = append(row, string(char))
		}
		fmt.Println()
		lines = append(lines, row)
		i++
	}

	println(day4part1(lines))
	println(day4part2(lines))
}
