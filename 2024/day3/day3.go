package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
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

func part1() int {
	res := 0
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day3/input1")
	check(err)

	reg, err := regexp.Compile(`mul\((\d{1,3},\d{1,3})\)`)
	check(err)

	for _, regResult := range reg.FindAllStringSubmatch(string(fileData), -1) {
		numbers := strings.Split(regResult[1], ",")
		number1, err := strconv.Atoi(numbers[0])
		check(err)
		number2, err := strconv.Atoi(numbers[1])
		check(err)

		res += number1 * number2
	}
	return res
}

func part2() int {
	res := 0
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day3/input2")
	check(err)

	reg, err := regexp.Compile(`mul\((\d{1,3},\d{1,3})\)|do\(\)|don't\(\)`)
	check(err)

	enabled := true

	for _, regResult := range reg.FindAllStringSubmatch(string(fileData), -1) {
		if regResult[0] == "do()" {
			enabled = true
		} else if regResult[0] == "don't()" {
			enabled = false
		}

		if !enabled || !strings.HasPrefix(regResult[0], "mul(") {
			continue
		}

		numbers := strings.Split(regResult[1], ",")
		number1, err := strconv.Atoi(numbers[0])
		check(err)
		number2, err := strconv.Atoi(numbers[1])
		check(err)

		res += number1 * number2
	}

	return res
}
