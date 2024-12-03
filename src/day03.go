package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func day3part1(content string) int {
	re := regexp.MustCompile(`(?m)mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(content, -1)
	var sum int
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}
	return sum
}

func day3part2(content string) int {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	isEnabled := true
	total := 0
	i := 0

	for i < len(content) {
		switch {
		case i+4 <= len(content) && doRegex.MatchString(content[i:i+4]):
			isEnabled = true
			i += 4
		case i+7 <= len(content) && dontRegex.MatchString(content[i:i+7]):
			isEnabled = false
			i += 7
		case i+12 <= len(content):
			slice := content[i : i+12]
			if matches := mulRegex.FindStringSubmatch(slice); matches != nil {
				if isEnabled {
					x, _ := strconv.Atoi(matches[1])
					y, _ := strconv.Atoi(matches[2])
					total += x * y
				}
				i += len(matches[0])
			} else {
				i++
			}
		default:
			i++
		}
	}
	return total
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

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content += line + "\n"
	}

	println(day3part1(content))
	println(day3part2(content))
}
