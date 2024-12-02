package main

import (
	"bufio"
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

// 1: 631; 2: 665
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

func part1() int {
	res := 0
	file, err := os.Open("/home/udz/advent_of_code/2024/day2/input1")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		report := strings.Split(scanner.Text(), " ")
		if isValidReport(report, -1) {
			res += 1
		}
	}

	return res
}

func part2() int {
	res := 0
	file, err := os.Open("/home/udz/advent_of_code/2024/day2/input2")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		report := strings.Split(scanner.Text(), " ")
		for j := 0; j < len(report); j++ {
			if isValidReport(report, j) {
				res += 1
				break
			}
		}
	}

	return res
}

func isValidReport(report []string, levelToSkip int) bool {
	var increasingRow bool
	var prev int
	processed := 0

	for j := 0; j < len(report); j++ {
		if j == levelToSkip {
			continue
		}

		number, err := strconv.Atoi(report[j])
		check(err)

		if processed == 0 {
			prev = number
			processed += 1
			continue
		}

		if processed == 1 {
			increasingRow = number > prev
		}

		if increasingRow && !(number-prev >= 1 && number-prev <= 3) {
			return false
		} else if !increasingRow && !(prev-number >= 1 && prev-number <= 3) {
			return false
		}

		prev = number
		processed += 1
	}

	return true
}
