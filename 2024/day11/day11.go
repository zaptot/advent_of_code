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

func part1() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day11/input1")
	check(err)

	data := strings.Split(string(fileData), " ")
	cache := make(map[int]map[int]int, 0)
	res := 0

	for i := 0; i < len(data); i++ {
		number, err := strconv.Atoi(data[i])
		check(err)

		res += helperWithCache(number, 75, cache)
	}

	return res
}

func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day11/input1")
	check(err)

	data := strings.Split(string(fileData), " ")
	cache := make(map[int]map[int]int, 0)
	res := 0

	for i := 0; i < len(data); i++ {
		number, err := strconv.Atoi(data[i])
		check(err)

		res += helperWithCache(number, 75, cache)
	}

	return res
}

func helperWithCache(num int, blinks int, cache map[int]map[int]int) int {
	if cache[num] == nil {
		cache[num] = make(map[int]int)
	}
	if cache[num][blinks] != 0 {
		return cache[num][blinks]
	}

	if blinks == 0 {
		return 1
	}

	if num == 0 {
		cache[num][blinks] = helperWithCache(1, blinks-1, cache)
		return cache[num][blinks]
	}

	size := digitSize(num)
	if size%2 == 0 {
		d := pow10half(size)
		num1 := num / d
		num2 := num % d
		cache[num][blinks] = helperWithCache(num1, blinks-1, cache) + helperWithCache(num2, blinks-1, cache)
		return cache[num][blinks]
	}

	cache[num][blinks] = helperWithCache(num*2024, blinks-1, cache)
	return cache[num][blinks]
}

func digitSize(digit int) int {
	res := 0

	for digit > 0 {
		res += 1
		digit /= 10
	}

	return res
}

func pow10half(size int) int {
	res := 1
	half := size / 2

	for size > half {
		res *= 10
		size -= 1
	}

	return res
}
