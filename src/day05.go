package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func day5FindIndex(array []int, target int) int {
	for i, v := range array {
		if v == target {
			return i
		}
	}
	return -1
}

func checkValidity(update []int, rules [][]int) (bool, int) {
	for i, updateNumber := range update {
		for _, rule := range rules {
			var index = day5FindIndex(update, rule[1])
			if updateNumber == rule[0] && index != -1 && i > index {
				return false, i
			}
		}
	}
	return true, -1
}

func day5part1(rules [][]int, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		var isUpdateValid, _ = checkValidity(update, rules)

		if isUpdateValid {
			var middle = update[(len(update) / 2)]
			sum += middle
		}
	}
	return sum
}

func day5part2(rules [][]int, updates [][]int) int {
	sum := 0
	for _, _update := range updates {
		//fmt.Printf("calculating %d/%d\n", i, len(updates))
		var isUpdateValid, _ = checkValidity(_update, rules)

		if !isUpdateValid {
			//fmt.Println("invalid update", _update)
			update := make([]int, len(_update))
			copy(update, _update)

			for valid, invalidIndex := checkValidity(update, rules); !valid; valid, invalidIndex = checkValidity(update, rules) {
				//fmt.Printf("%v, %d, %d, %v\n", valid, invalidIndex, update[invalidIndex], update)
				pop := update[invalidIndex]
				update = append(update[:invalidIndex], update[invalidIndex+1:]...)
				update = append([]int{pop}, update...)
			}

			var middle = update[(len(update) / 2)]
			sum += middle
		}
	}
	return sum
}

func main() {
	filePath := filepath.Join("data", "2024", "day5.txt")

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
	// todo: create map from it
	var rules [][]int
	var updates [][]int

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			var rule = strings.Split(line, "|")
			var number, _ = strconv.Atoi(rule[0])
			var numberRule, _ = strconv.Atoi(rule[1])

			rules = append(rules, []int{number, numberRule})
		}
		if strings.Contains(line, ",") {
			var numbers = strings.Split(line, ",")
			var step []int
			for _, number := range numbers {
				var n, _ = strconv.Atoi(number)
				step = append(step, n)
			}

			updates = append(updates, step)
		}
	}

	println(day5part1(rules, updates))
	println(day5part2(rules, updates))
}
