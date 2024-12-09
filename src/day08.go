package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type Coords struct {
	x, y int
}

func day8part1(grid [][]rune) int {
	count := 0
	antennas := make(map[rune][]Coords)
	gridCopy := make([][]rune, len(grid))
	for i := range grid {
		gridCopy[i] = make([]rune, len(grid[i]))
		copy(gridCopy[i], grid[i])

		for j := range grid[i] {
			if grid[i][j] != '.' {
				_, ok := antennas[grid[i][j]]
				if !ok {
					antennas[grid[i][j]] = make([]Coords, 0)
				}
				antennas[grid[i][j]] = append(antennas[grid[i][j]], Coords{i, j})
			}
		}
	}

	for _, antenna := range antennas {
		for _, antennaCoords := range antenna {
			for _, otherAntennaCoords := range antenna {
				if antennaCoords.x == otherAntennaCoords.x && antennaCoords.y == otherAntennaCoords.y {
					continue
				}
				dx, dy := antennaCoords.x-otherAntennaCoords.x, antennaCoords.y-otherAntennaCoords.y
				newX, newY := antennaCoords.x+dx, antennaCoords.y+dy

				if newX >= 0 && newY >= 0 && newX < len(gridCopy) && newY < len(gridCopy[0]) {
					if gridCopy[newX][newY] != '#' {
						gridCopy[newX][newY] = '#'
						count++
					}
				}
			}
		}
	}
	return count
}

func day8part2(grid [][]rune) int {
	count := 0
	antennas := make(map[rune][]Coords)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '.' {
				_, ok := antennas[grid[i][j]]
				if !ok {
					antennas[grid[i][j]] = make([]Coords, 0)
				}
				antennas[grid[i][j]] = append(antennas[grid[i][j]], Coords{i, j})
			}
		}
	}

	for _, antenna := range antennas {
		for _, antennaCoords := range antenna {
			for _, otherAntennaCoords := range antenna {
				if antennaCoords.x == otherAntennaCoords.x && antennaCoords.y == otherAntennaCoords.y {
					continue
				}
				dx, dy := antennaCoords.x-otherAntennaCoords.x, antennaCoords.y-otherAntennaCoords.y
				newX, newY := antennaCoords.x+dx, antennaCoords.y+dy

				for newX >= 0 && newY >= 0 && newX < len(grid) && newY < len(grid[0]) {
					if grid[newX][newY] != '#' {
						grid[newX][newY] = '#'
					}
					newX += dx
					newY += dy
				}
			}
		}
	}
	for _, row := range grid {
		for _, col := range row {
			if col != '.' {
				count++
			}
		}
	}
	return count
}

func main() {
	filePath := filepath.Join("data", "2024", "day8.txt")

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
	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		nodes := []rune(line)
		grid = append(grid, nodes)
	}

	println(day8part1(grid))
	println(day8part2(grid))
}
