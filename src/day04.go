package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func day4checkForXMAS(grid [][]string, x int, y int) int {
	directions := [][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1}, // Vertical and horizontal
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1}, // Diagonal
	}
	count := 0

	for _, direction := range directions {
		dx, dy := direction[0], direction[1]
		if x+3*dx >= 0 && x+3*dx < len(grid) && y+3*dy >= 0 && y+3*dy < len(grid[0]) &&
			grid[x][y] == "X" && grid[x+dx][y+dy] == "M" && grid[x+2*dx][y+2*dy] == "A" && grid[x+3*dx][y+3*dy] == "S" {
			count++
		}
	}

	return count
}

func day4checkForMAS(grid [][]string, x int, y int) int {
	count := 0
	if x-1 < 0 || x+1 >= len(grid) || y-1 < 0 || y+1 >= len(grid[0]) {
		return 0
	}

	if grid[x-1][y-1] == "M" && grid[x+1][y+1] == "S" && grid[x+1][y-1] == "M" && grid[x-1][y+1] == "S" ||
		grid[x-1][y+1] == "M" && grid[x+1][y-1] == "S" && grid[x+1][y+1] == "M" && grid[x-1][y-1] == "S" {
		count++
	}
	if grid[x-1][y-1] == "M" && grid[x+1][y+1] == "S" && grid[x-1][y+1] == "M" && grid[x+1][y-1] == "S" ||
		grid[x+1][y-1] == "M" && grid[x+-1][y+1] == "S" && grid[x+1][y+1] == "M" && grid[x-1][y-1] == "S" {
		count++
	}

	return count
}

func day4part1(grid [][]string) int {
	count := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == "X" {
				count += day4checkForXMAS(grid, i, j)
			}
		}
	}

	return count
}

func day4part2(grid [][]string) int {
	count := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == "A" {
				count += day4checkForMAS(grid, i, j)
			}
		}
	}

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
	var grid [][]string

	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	println(day4part1(grid))
	println(day4part2(grid))
}
