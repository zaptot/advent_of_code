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
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day12/input1")
	check(err)

	fileRows := strings.Split(string(fileData), "\n")
	visited := make(map[int]map[int]bool)
	res := 0

	for i := 0; i < len(fileRows); i++ {
		for j := 0; j < len(fileRows[i]); j++ {
			if visited[i][j] {
				continue
			}

			perimeter, area := helper(fileRows, i, j, visited, getFieldPerimeter)
			mult := perimeter * area
			res += mult
		}
	}

	return res
}

func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day12/input1")
	check(err)

	fileRows := strings.Split(string(fileData), "\n")
	visited := make(map[int]map[int]bool)
	res := 0

	for i := 0; i < len(fileRows); i++ {
		for j := 0; j < len(fileRows[i]); j++ {
			if visited[i][j] {
				continue
			}

			sides, area := helper(fileRows, i, j, visited, getFieldSlides)
			mult := sides * area
			res += mult
		}
	}

	return res
}

func helper(
	fileRows []string,
	i int,
	j int,
	visited map[int]map[int]bool,
	caculator func([]string, int, int) int,
) (int, int) {
	if visited[i][j] {
		return 0, 0
	}

	if visited[i] == nil {
		visited[i] = make(map[int]bool)
	}

	visited[i][j] = true
	directions := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	perimeter := caculator(fileRows, i, j)
	area := 1

	for di := 0; di < len(directions); di++ {
		nextI := i + directions[di][0]
		nextJ := j + directions[di][1]
		if nextI < 0 || nextJ < 0 || nextI >= len(fileRows) || nextJ >= len(fileRows[0]) ||
			fileRows[i][j] != fileRows[nextI][nextJ] {
			continue
		}

		subP, subA := helper(fileRows, nextI, nextJ, visited, caculator)
		perimeter += subP
		area += subA
	}

	return perimeter, area
}

func getFieldPerimeter(fileRows []string, i int, j int) int {
	res := 0
	currentFlower := fileRows[i][j]

	if i == 0 || fileRows[i-1][j] != currentFlower {
		res += 1
	}

	if j == 0 || fileRows[i][j-1] != currentFlower {
		res += 1
	}

	if i == len(fileRows)-1 || fileRows[i+1][j] != currentFlower {
		res += 1
	}

	if j == len(fileRows[i])-1 || fileRows[i][j+1] != currentFlower {
		res += 1
	}

	return res
}

func getFieldSlides(fileRows []string, i int, j int) int {
	res := 0

	currentFlower := fileRows[i][j]

	// bot
	if ((i < len(fileRows)-1 && fileRows[i+1][j] != currentFlower) || i == len(fileRows)-1) &&
		(j == 0 || fileRows[i][j-1] != currentFlower ||
			(i < len(fileRows)-1 && fileRows[i+1][j-1] == currentFlower)) {
		res += 1
	}

	//top
	if ((i > 0 && fileRows[i-1][j] != currentFlower) || i == 0) &&
		(j == 0 || fileRows[i][j-1] != currentFlower ||
			(i > 0 && fileRows[i-1][j-1] == currentFlower)) {
		res += 1
	}

	//right
	if ((j < len(fileRows[i])-1 && fileRows[i][j+1] != currentFlower) || j == len(fileRows[i])-1) &&
		(i == 0 || fileRows[i-1][j] != currentFlower ||
			(j < len(fileRows[i])-1 && fileRows[i-1][j+1] == currentFlower)) {
		res += 1
	}

	//left
	if ((j > 0 && fileRows[i][j-1] != currentFlower) || j == 0) &&
		(i == 0 || fileRows[i-1][j] != currentFlower ||
			(j > 0 && fileRows[i-1][j-1] == currentFlower)) {
		res += 1
	}

	return res
}
