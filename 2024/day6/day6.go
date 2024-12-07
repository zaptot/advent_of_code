package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const LEFT = 0
const RIGHT = 1
const TOP = 2
const BOT = 3

var DIRECTIONS_CHANGES = map[int]int{
	TOP:   RIGHT,
	RIGHT: BOT,
	BOT:   LEFT,
	LEFT:  TOP,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 3, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1()
		fmt.Println("Output:", ans)
	} else if part == 2 {
		startTime := time.Now().UnixMilli()
		ans := part2()
		fmt.Println("Output:", ans)
		endtTime := time.Now().UnixMilli()

		fmt.Printf("time: %d ms \n", endtTime-startTime)
	} else {
		startTime := time.Now().UnixMilli()
		ans := part2Optimized()
		fmt.Println("Output:", ans)
		endtTime := time.Now().UnixMilli()

		fmt.Printf("time: %d ms \n", endtTime-startTime)
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

func part2Optimized() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day6/input1")
	check(err)

	board := strings.Split(string(fileData), "\n")
	cache := preprocessNearestObstacles(board)
	currentPosition := findStartPosition(board)
	visited := make(map[[2]int]map[int]bool)
	moveWithoutObstacles(board, currentPosition, visited)
	res := 0

	for obstaclePosition := range visited {
		cycle := optimizedIsCycle(currentPosition, obstaclePosition, cache)

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

func optimizedIsCycle(
	startPosition [2]int,
	obstaclePosition [2]int,
	cache [][][4]int,
) bool {
	pos := [2]int{startPosition[0], startPosition[1]}
	currentDirection := TOP
	visited := make(map[[2]int]map[int]bool)
	cycle := false

	for {
		if visited[pos][currentDirection] {
			cycle = true
			break
		}
		if visited[pos] == nil {
			visited[pos] = make(map[int]bool)
		}

		visited[pos][currentDirection] = true
		nextPos := [2]int{-1, -1}

		// check added obstacle
		if currentDirection == RIGHT && pos[0] == obstaclePosition[0] && pos[1] < obstaclePosition[1] {
			nextPos = [2]int{pos[0], obstaclePosition[1] - 1}
		} else if currentDirection == LEFT && pos[0] == obstaclePosition[0] && pos[1] > obstaclePosition[1] {
			nextPos = [2]int{pos[0], obstaclePosition[1] + 1}
		} else if currentDirection == TOP && pos[1] == obstaclePosition[1] && pos[0] > obstaclePosition[0] {
			nextPos = [2]int{obstaclePosition[0] + 1, pos[1]}
		} else if currentDirection == BOT && pos[1] == obstaclePosition[1] && pos[0] < obstaclePosition[0] {
			nextPos = [2]int{obstaclePosition[0] - 1, pos[1]}
		}

		if cache[pos[0]][pos[1]][currentDirection] == -1 && nextPos == [2]int{-1, -1} {
			break
		}

		// no obstacles in current direction
		if cache[pos[0]][pos[1]][currentDirection] == -1 {
			//do nothing
		} else if currentDirection == RIGHT {
			if nextPos == [2]int{-1, -1} || nextPos[1] > cache[pos[0]][pos[1]][currentDirection]-1 {
				nextPos = [2]int{pos[0], cache[pos[0]][pos[1]][currentDirection] - 1}
			}
		} else if currentDirection == LEFT {
			if nextPos == [2]int{-1, -1} || nextPos[1] < cache[pos[0]][pos[1]][currentDirection]+1 {
				nextPos = [2]int{pos[0], cache[pos[0]][pos[1]][currentDirection] + 1}
			}
		} else if currentDirection == TOP {
			if nextPos == [2]int{-1, -1} || nextPos[0] < cache[pos[0]][pos[1]][currentDirection]+1 {
				nextPos = [2]int{cache[pos[0]][pos[1]][currentDirection] + 1, pos[1]}
			}
		} else if currentDirection == BOT {
			if nextPos == [2]int{-1, -1} || nextPos[0] > cache[pos[0]][pos[1]][currentDirection]-1 {
				nextPos = [2]int{cache[pos[0]][pos[1]][currentDirection] - 1, pos[1]}
			}
		} else {
			panic("wrong next obstacle condition")
		}

		pos = nextPos
		currentDirection = DIRECTIONS_CHANGES[currentDirection]
	}

	return cycle
}

func preprocessNearestObstacles(board []string) [][][4]int {
	cache := make([][][4]int, len(board))

	for i := 0; i < len(board); i++ {
		cache[i] = make([][4]int, len(board[i]))

		for j := 0; j < len(board[i]); j++ {
			cache[i][j] = [4]int{-1, -1, -1, -1}
		}
	}

	for i := 0; i < len(board); i++ {
		prevObstacleIdx := -1
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '#' {
				prevObstacleIdx = j
			}

			cache[i][j][LEFT] = prevObstacleIdx
		}
	}

	for j := 0; j < len(board[0]); j++ {
		prevObstacleIdx := -1
		for i := 0; i < len(board); i++ {
			if board[i][j] == '#' {
				prevObstacleIdx = i
			}

			cache[i][j][TOP] = prevObstacleIdx
		}
	}

	for i := 0; i < len(board)-1; i++ {
		prevObstacleIdx := -1
		for j := len(board[i]) - 1; j >= 0; j-- {
			if board[i][j] == '#' {
				prevObstacleIdx = j
			}

			cache[i][j][RIGHT] = prevObstacleIdx
		}
	}

	for j := 0; j < len(board[0]); j++ {
		prevObstacleIdx := -1
		for i := len(board) - 1; i >= 0; i-- {
			if board[i][j] == '#' {
				prevObstacleIdx = i
			}

			cache[i][j][BOT] = prevObstacleIdx
		}
	}

	return cache
}
