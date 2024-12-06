package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

var directions = []coords{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type coords struct {
	x, y int
}

func isValidIndex(x, y, maxX, maxY int) bool {
	return x < maxX && y < maxY && x > 0 && y > 0
}

func hasCycle(grid [][]rune, start coords) bool {
	visited := make(map[coords]coords)
	curr := start
	currDir := 0
	maxX, maxY := len(grid), len(grid[0])

	move := func(c coords, dir coords) coords {
		return coords{x: c.x + dir.x, y: c.y + dir.y}
	}

	for isValidIndex(curr.x, curr.y, maxX, maxY) {
		if visited[curr] == directions[currDir] {
			return true
		}

		visited[curr] = directions[currDir]

		if grid[curr.x][curr.y] == '#' {
			curr = move(curr, coords{-directions[currDir].x, -directions[currDir].y})
			currDir = (currDir + 1) % len(directions)
		} else {
			curr = move(curr, directions[currDir])
		}
	}

	return false
}

func day6part1(grid [][]rune, start coords) int {
	maxX, maxY := len(grid), len(grid[0])
	visited := make(map[coords]bool)
	index := start
	currDir := 0

	for isValidIndex(index.x, index.y, maxX, maxY) {
		visited[index] = true
		dx, dy := directions[currDir].x, directions[currDir].y
		if grid[index.x+dx][index.y+dy] == '#' {
			currDir = (currDir + 1) % len(directions)
		} else {
			index.x, index.y = index.x+dx, index.y+dy
		}
	}

	count := 0
	for range visited {
		count++
	}
	return count + 1
}

func day6part2(grid [][]rune, start coords) int {
	maxX, maxY := len(grid), len(grid[0])
	visited := make(map[coords]bool)
	index := start
	currDir := 0

	for isValidIndex(index.x, index.y, maxX, maxY) {
		visited[index] = true

		dx, dy := directions[currDir].x, directions[currDir].y
		if isValidIndex(index.x+dx, index.y+dy, maxX, maxY) && grid[index.x+dx][index.y+dy] == '#' {
			currDir = (currDir + 1) % len(directions)
		} else {
			index.x, index.y = index.x+dx, index.y+dy
		}
	}

	obstructionCount := 0
	for pos := range visited {
		if pos == start || grid[pos.x][pos.y] == '#' {
			continue
		}
		grid[pos.x][pos.y] = '#'
		if hasCycle(grid, start) {
			obstructionCount++
		}
		grid[pos.x][pos.y] = '.'
	}

	return obstructionCount
}

func main() {
	filePath := filepath.Join("data", "2024", "day6.txt")

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
	var start coords

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		if c := slices.Index(row, '^'); c != -1 {
			start = coords{x: len(grid), y: c}
			row[c] = '.' // Replace starting point
		}
		grid = append(grid, row)
	}

	println(day6part1(grid, start))
	println(day6part2(grid, start))
}
