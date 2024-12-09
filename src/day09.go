package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

type day9Block struct {
	value int
	size  int
}

func day9part1(disk string) int {
	code := []rune(disk)
	var interpretation []day9Block
	for id, i := 0, 0; i < len(code); i++ {
		repeatCount, _ := strconv.Atoi(string(code[i]))
		if i%2 == 0 {
			for _ = range repeatCount {
				interpretation = append(interpretation, day9Block{id, repeatCount})
			}
		} else {
			for _ = range repeatCount {
				interpretation = append(interpretation, day9Block{-1, repeatCount})
			}
			id++
		}

	}

	fragmentation := make([]day9Block, len(interpretation))
	copy(fragmentation, interpretation)

	for {
		freeIndex := slices.IndexFunc(fragmentation, func(b day9Block) bool { return b.value == -1 })
		var toFragmentIndex int
		for i := len(fragmentation) - 1; i >= 0; i-- {
			if fragmentation[i].value != -1 {
				toFragmentIndex = i
				break
			}
		}
		if freeIndex < toFragmentIndex {
			fragmentation[freeIndex], fragmentation[toFragmentIndex] = fragmentation[toFragmentIndex], fragmentation[freeIndex]
		} else {
			break
		}
	}

	checksum := 0
	for i, char := range fragmentation {
		if char.value == -1 {
			continue
		}
		checksum += char.value * i
	}
	return checksum
}

func day9part2(disk string) int {
	code := []rune(disk)
	var interpretation []day9Block
	for id, i := 0, 0; i < len(code); i++ {
		repeatCount, _ := strconv.Atoi(string(code[i]))
		if i%2 == 0 {
			for _ = range repeatCount {
				interpretation = append(interpretation, day9Block{id, repeatCount})
			}
		} else {
			for _ = range repeatCount {
				interpretation = append(interpretation, day9Block{-1, repeatCount})
			}
			id++
		}

	}

	fragmentation := make([]day9Block, len(interpretation))
	copy(fragmentation, interpretation)

	blockIndexToMove := len(fragmentation) - 1
	freeSpaceIndex := slices.IndexFunc(fragmentation, func(b day9Block) bool { return b.value == -1 })

	nextVal := fragmentation[blockIndexToMove].value

	for nextVal >= 0 {
		currBlock := fragmentation[blockIndexToMove]
		freeSpaceIndex = 0
		found := false
	outer:
		for i := freeSpaceIndex; i < len(fragmentation)-1; i++ {
			if fragmentation[i].value == -1 {
				for j := 0; j < currBlock.size; j++ {
					if (i+j) > len(fragmentation)-1 || fragmentation[i+j].value != -1 {
						continue outer
					}
				}
				found = true
				freeSpaceIndex = i
				break outer
			}
		}
		if found && freeSpaceIndex < blockIndexToMove {
			for i := 0; i < currBlock.size; i++ {
				fragmentation[blockIndexToMove-i], fragmentation[freeSpaceIndex+i] = fragmentation[freeSpaceIndex+i], fragmentation[blockIndexToMove-i]
			}
		}
		nextVal -= 1
		blockIndexToMove -= currBlock.size
		for i := blockIndexToMove; i >= 0; i-- {
			if fragmentation[i].value != -1 && fragmentation[i].value == nextVal {
				blockIndexToMove = i
				break
			}
		}
	}

	checksum := 0
	for i, char := range fragmentation {
		if char.value == -1 {
			continue
		}
		checksum += char.value * i
	}
	return checksum
}

func main() {
	filePath := filepath.Join("data", "2024", "day9.txt")

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
	disk := string(input)

	fmt.Println(day9part1(disk))
	fmt.Println(day9part2(disk))
}
