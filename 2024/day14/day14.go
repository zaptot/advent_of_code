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

type Robot struct {
	x  int
	y  int
	dx int
	dy int
}

func (r *Robot) Move(stepsCount int, boardWidth int, boardHeight int) {
	newX := (r.x + (stepsCount * r.dx)) % boardWidth
	newY := (r.y + (stepsCount * r.dy)) % boardHeight
	if newX < 0 {
		newX = boardWidth + newX
	}

	if newY < 0 {
		newY = boardHeight + newY
	}

	r.x = newX
	r.y = newY
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
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day14/input1")
	check(err)

	fileRows := strings.Split(string(fileData), "\n")
	robots := make([]*Robot, len(fileRows))
	steps := 100
	boardWidth := 101
	boardHeight := 103

	for i, robotString := range fileRows {
		robots[i] = parseRobot(robotString)
	}

	for _, robot := range robots {
		robot.Move(steps, boardWidth, boardHeight)

	}

	return calculateSafeScore(robots, boardWidth, boardHeight)
}

func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day14/input1")
	check(err)

	fileRows := strings.Split(string(fileData), "\n")
	robots := make([]*Robot, len(fileRows))
	boardWidth := 101
	boardHeight := 103

	for i, robotString := range fileRows {
		robots[i] = parseRobot(robotString)
		robots[i].Move(7753, boardWidth, boardHeight)
	}

	writePositionsIntoFile(robots, boardWidth, boardHeight, 7753)
	return 7753
}

func parseRobot(robotString string) *Robot {
	positionAndSpeed := strings.Split(robotString, " ")
	position := strings.Split(strings.SplitAfter(positionAndSpeed[0], "p=")[1], ",")
	speed := strings.Split(strings.SplitAfter(positionAndSpeed[1], "v=")[1], ",")

	positionX, err := strconv.Atoi(position[0])
	check(err)

	positionY, err := strconv.Atoi(position[1])
	check(err)

	speedX, err := strconv.Atoi(speed[0])
	check(err)

	speedY, err := strconv.Atoi(speed[1])
	check(err)

	return &Robot{x: positionX, y: positionY, dx: speedX, dy: speedY}
}

func calculateSafeScore(robots []*Robot, boardWidth int, boardHeight int) int {
	first := 0
	second := 0
	third := 0
	fourth := 0

	for _, robot := range robots {
		if robot.x < boardWidth/2 && robot.y < boardHeight/2 {
			first += 1
		} else if robot.x > boardWidth/2 && robot.y < boardHeight/2 {
			second += 1
		} else if robot.x < boardWidth/2 && robot.y > boardHeight/2 {
			third += 1
		} else if robot.x > boardWidth/2 && robot.y > boardHeight/2 {
			fourth += 1
		}
	}

	return first * second * third * fourth
}

func writePositionsIntoFile(robots []*Robot, boardWidth int, boardHeight int, id int) {
	positions := make([][]int, boardHeight)

	for _, robot := range robots {
		if positions[robot.y] == nil {
			positions[robot.y] = make([]int, boardWidth)
		}

		positions[robot.y][robot.x] += 1
	}

	fmt.Println(positions)

	rowsToWrite := make([]byte, boardHeight*(boardWidth+1))

	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(positions[0]); j++ {
			pos := (i * (boardWidth + 1)) + j
			rowsToWrite[pos] = ' '

			if positions[i] != nil && positions[i][j] > 0 {
				rowsToWrite[pos] = '#'
			}
		}
		rowsToWrite[(i*(boardWidth+1))+boardWidth] = '\n'
	}

	fileName := "/home/udz/advent_of_code/2024/day14/res/" + strconv.Itoa(id)
	err := os.WriteFile(fileName, rowsToWrite, 0644)
	check(err)
}
