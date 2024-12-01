package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
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
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day1/input1")
	check(err)

	fileRows := strings.Split(string(fileData), "\n")
	numbers := [][]int{make([]int, len(fileRows)), make([]int, len(fileRows))}

	for i := 0; i < len(fileRows); i++ {
		parsedNumbers := strings.Split(fileRows[i], "   ")
		firstNumber, err := strconv.Atoi(parsedNumbers[0])
		check(err)
		secondNumber, err := strconv.Atoi(parsedNumbers[1])
		check(err)

		numbers[0][i] = firstNumber
		numbers[1][i] = secondNumber
	}

	slices.Sort(numbers[0])
	slices.Sort(numbers[1])

	res := 0

	for i := 0; i < len(numbers[0]); i++ {
		res += int(math.Abs(float64(numbers[1][i] - numbers[0][i])))
	}

	return res
}

func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day1/input2")
	check(err)

	fileRows := strings.Split(string(fileData), "\n")
	firstColNumbers := make([]int, len(fileRows))
	secondColCounts := make(map[int]int)

	for i := 0; i < len(fileRows); i++ {
		parsedNumbers := strings.Split(fileRows[i], "   ")
		firstNumber, err := strconv.Atoi(parsedNumbers[0])
		check(err)
		secondNumber, err := strconv.Atoi(parsedNumbers[1])
		check(err)
		firstColNumbers[i] = firstNumber

		secondColCounts[secondNumber] += 1
	}

	res := 0

	for i := 0; i < len(fileRows); i++ {
		res += firstColNumbers[i] * secondColCounts[firstColNumbers[i]]
	}

	return res
}
