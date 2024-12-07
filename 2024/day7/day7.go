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

// 28730327770375
func part1() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day7/input1")
	check(err)

	fileRows := strings.Split(string(fileData), "\n")
	res := 0

	for i := 0; i < len(fileRows); i++ {
		row := strings.Split(fileRows[i], ": ")
		rowResult, err := strconv.Atoi(row[0])
		check(err)
		rowNumbers := strings.Split(row[1], " ")
		numbers := make([]int, len(rowNumbers))

		for idx, number := range rowNumbers {
			parsedNumber, err := strconv.Atoi(number)
			check(err)

			numbers[idx] = parsedNumber
		}

		if canBeCombinedToResult(numbers, rowResult, []func(int, int) int{addition, miltiplication}) {
			res += rowResult
		}
	}

	return res
}

// 424977609625985
func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day7/input2")
	check(err)

	fileRows := strings.Split(string(fileData), "\n")
	res := 0

	fmt.Println(concatenation(1, 100))
	for i := 0; i < len(fileRows); i++ {
		row := strings.Split(fileRows[i], ": ")
		rowResult, err := strconv.Atoi(row[0])
		check(err)
		rowNumbers := strings.Split(row[1], " ")
		numbers := make([]int, len(rowNumbers))

		for idx, number := range rowNumbers {
			parsedNumber, err := strconv.Atoi(number)
			check(err)

			numbers[idx] = parsedNumber
		}

		if canBeCombinedToResult(numbers, rowResult, []func(int, int) int{addition, miltiplication, concatenation}) {
			res += rowResult
		} else {
			// fmt.Println(rowResult)
		}
	}

	return res
}

func canBeCombinedToResult(numbers []int, result int, fs []func(int, int) int) bool {
	return helper(numbers, result, numbers[0], 1, fs)
}

func helper(numbers []int, neededResult int, currentResult int, currentIdx int, fs []func(int, int) int) bool {
	if currentResult > neededResult {
		return false
	}

	if currentIdx >= len(numbers) {
		return neededResult == currentResult
	}

	res := false

	for _, fn := range fs {
		if res {
			break
		}

		res = res || helper(numbers, neededResult, fn(currentResult, numbers[currentIdx]), currentIdx+1, fs)
	}

	return res
}

func addition(a, b int) int {
	return a + b
}

func miltiplication(a, b int) int {
	return a * b
}

func concatenation(a, b int) int {
	res := a
	reversedB := 1

	for b > 0 {
		reversedB *= 10
		reversedB += b % 10
		b /= 10
	}

	for reversedB > 1 {
		res *= 10
		res += reversedB % 10
		reversedB /= 10
	}

	return res
}
