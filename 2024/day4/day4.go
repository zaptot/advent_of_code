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
	res := 0
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day4/input1")
	check(err)

	board := strings.Split(string(fileData), "\n")
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			for k := 0; k < len(directions); k++ {
				if isXmas(board, i, j, directions[k]) {
					res += 1
				}
			}
		}
	}
	return res
}

func part2() int {
	res := 0
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day4/input2")
	check(err)

	board := strings.Split(string(fileData), "\n")

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			if isMas(board, i, j) {
				res += 1
			}
		}
	}

	return res
}

func isXmas(board []string, row int, col int, direction [2]int) bool {
	// fmt.Printf("row %d, col %d \n", row, col)
	word := "XMAS"

	for i := 0; i < 4; i++ {
		if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) {
			return false
		}
		if board[row][col] != word[i] {
			return false
		}

		row += direction[0]
		col += direction[1]
	}

	return true
}

func isMas(board []string, row int, col int) bool {
	if row < 1 || col < 1 || row+1 >= len(board) || col+1 > len(board[0]) || board[row][col] != 'A' {
		return false
	}

	directions := [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	ms := 0
	ss := 0

	for i := 0; i < len(directions); i++ {
		if board[row+directions[i][0]][col+directions[i][1]] == 'M' {
			ms += 1
		} else if board[row+directions[i][0]][col+directions[i][1]] == 'S' {
			ss += 1
		}
	}

	if ms != 2 || ss != 2 || board[row-1][col-1] == board[row+1][col+1] {
		return false
	}

	return true
}
