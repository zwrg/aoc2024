package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func main() {
	filePath := filepath.Join("data", "2024", "day3.txt")

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	content := string(data)
	total1, total2 := 0, 0
	isEnabled := true

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|(do\(\))|(don't\(\))`)

	matches := re.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if match[4] != "" {
			isEnabled = false
		} else if match[3] != "" {
			isEnabled = true
		} else {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			total1 += x * y
			if isEnabled {
				total2 += x * y
			}
		}
	}

	fmt.Println(total1)
	fmt.Println(total2)
}
