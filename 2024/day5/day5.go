package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 2, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1()
		fmt.Println("Output:", ans)
	} else {
		ans := part2()
		fmt.Println("Output:", ans)
	}
}

// 4766
func part1() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day5/input1")
	check(err)

	fileParts := strings.Split(string(fileData), "\n\n")

	rules := parseRules(strings.Split(fileParts[0], "\n"))
	pages := parsePages(strings.Split(fileParts[1], "\n"))
	res := 0

	for i := 0; i < len(pages); i++ {
		page := pages[i]

		if isRowValid(page, rules) {
			midElIdx := (len(page) / 2)

			res += page[midElIdx]
		}
	}

	return res
}

func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day5/input2")
	check(err)

	fileParts := strings.Split(string(fileData), "\n\n")
	rules := parseRules(strings.Split(fileParts[0], "\n"))
	pages := parsePages(strings.Split(fileParts[1], "\n"))
	res := 0

	for i := 0; i < len(pages); i++ {
		page := pages[i]
		if !isRowValid(page, rules) {
			makeRowValid(page, rules)

			if isRowValid(page, rules) {
				midElIdx := (len(page) / 2)

				res += page[midElIdx]
			}
		}
	}

	return res
}

func parseRules(rulesRows []string) map[int]map[int]bool {
	rules := make(map[int]map[int]bool)

	for i := 0; i < len(rulesRows); i++ {
		rule := strings.Split(rulesRows[i], "|")
		firstNumber, err := strconv.Atoi(rule[0])
		check(err)
		secondNumber, err := strconv.Atoi(rule[1])
		check(err)

		if rules[firstNumber] == nil {
			rules[firstNumber] = make(map[int]bool)
		}

		rules[firstNumber][secondNumber] = true
	}

	return rules
}

func parsePages(pages []string) [][]int {
	res := make([][]int, len(pages))

	for i := 0; i < len(pages); i++ {
		pageNumbers := strings.Split(pages[i], ",")
		res[i] = make([]int, len(pageNumbers))

		for j := 0; j < len(pageNumbers); j++ {
			number, err := strconv.Atoi(pageNumbers[j])
			check(err)

			res[i][j] = number
		}
	}

	return res
}

func isRowValid(rowNumbers []int, rules map[int]map[int]bool) bool {
	rowMap := make(map[int]int)

	for i := 0; i < len(rowNumbers); i++ {
		rowMap[rowNumbers[i]] = i
	}

	for i := 0; i < len(rowNumbers); i++ {
		if rules[rowNumbers[i]] == nil {
			continue
		}

		for secondNumber := range rules[rowNumbers[i]] {
			secondIndex, ok := rowMap[secondNumber]
			if !ok {
				continue
			}

			if secondIndex <= i {
				return false
			}
		}
	}
	return true
}

func makeRowValid(page []int, rules map[int]map[int]bool) {
	for i := 0; i < len(page); i++ {
		for j := i + 1; j < len(page); j++ {
			firstNumber := page[i]
			secondNumber := page[j]

			if rules[firstNumber] == nil && rules[secondNumber] == nil {
				continue
			}

			if rules[firstNumber][secondNumber] {
				continue
			}

			if rules[secondNumber][firstNumber] {
				page[i] = secondNumber
				page[j] = firstNumber
			}
		}
	}
}
