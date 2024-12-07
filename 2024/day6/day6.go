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

// 5461
func part1() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day6/input1")
	check(err)

	board := strings.Split(string(fileData), "\n")
	currentPosition := findStartPosition(board)
	visited := make(map[[2]int]map[int]bool)

	visitedCount, _ := moveWithoutObstacles(board, currentPosition, visited)

	return visitedCount
}

// 1836
func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day6/input1")
	check(err)

	board := strings.Split(string(fileData), "\n")
	currentPosition := findStartPosition(board)
	visited := make(map[[2]int]map[int]bool)
	moveWithoutObstacles(board, currentPosition, visited)
	res := 0

	for obstaclePosition := range visited {
		_, cycle := move(board, currentPosition, make(map[[2]int]map[int]bool), obstaclePosition)

		if cycle {
			res += 1
		}
	}

	return res
}

func findStartPosition(board []string) [2]int {
	character := '^'

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == byte(character) {
				return [2]int{i, j}
			}
		}
	}

	panic("position not found")
}

func moveWithoutObstacles(board []string, startPosition [2]int, visited map[[2]int]map[int]bool) (int, bool) {
	return move(board, startPosition, visited, [2]int{-1, -1})
}

func move(board []string, startPosition [2]int, visited map[[2]int]map[int]bool, obstaclePosition [2]int) (int, bool) {
	pos := [2]int{startPosition[0], startPosition[1]}
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	currentDirection := 0
	visited[startPosition] = map[int]bool{currentDirection: true}

	for {
		if pos[0] == 0 || pos[0] == len(board)-1 ||
			pos[1] == 0 || pos[1] == len(board[1])-1 {
			break
		}
		direction := directions[currentDirection%len(directions)]

		for {
			if pos[0] == 0 || pos[1] == 0 || pos[0] == len(board)-1 || pos[1] == len(board[0])-1 {
				break
			}

			nextPos := [2]int{pos[0] + direction[0], pos[1] + direction[1]}

			if board[nextPos[0]][nextPos[1]] == '#' || nextPos == obstaclePosition {
				break
			}
			if visited[nextPos] == nil {
				visited[nextPos] = make(map[int]bool)
			}
			if visited[nextPos][currentDirection%len(directions)] {
				return -1, true
			}

			visited[nextPos][currentDirection%len(directions)] = true

			pos = nextPos
		}
		currentDirection += 1
	}

	return len(visited), false
}
