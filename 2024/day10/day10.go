package main

import (
	"flag"
	"fmt"
	"os"
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
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day10/input1")
	check(err)

	board := strings.Split(string(fileData), "\n")
	res := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '0' {
				continue
			}

			visited := make(map[[2]int]int)
			helper(board, i, j, visited)

			res += len(visited)
		}
	}
	return res
}

func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day10/input1")
	check(err)

	board := strings.Split(string(fileData), "\n")
	res := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '0' {
				continue
			}

			visited := make(map[[2]int]int)
			helper(board, i, j, visited)

			for _, count := range visited {
				res += count
			}
		}
	}
	return res
}

func helper(board []string, i int, j int, visited map[[2]int]int) {
	if board[i][j] == '9' {
		visited[[2]int{i, j}] += 1
		return
	}

	directions := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for d := 0; d < len(directions); d++ {
		di := i + directions[d][0]
		dj := j + directions[d][1]
		if di < 0 || dj < 0 || di >= len(board[0]) || dj >= len(board[1]) ||
			board[di][dj] != board[i][j]+1 {
			continue
		}

		helper(board, di, dj, visited)
	}
}
