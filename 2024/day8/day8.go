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
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day8/input1")
	check(err)

	board := strings.Split(string(fileData), "\n")
	groupedAntennas := groupAllAntennas(board)
	createdAntinodes := make(map[[2]int]bool)
	// fmt.Println(groupedAntennas)

	for _, groupedPositions := range groupedAntennas {
		for i := 0; i < len(groupedPositions); i++ {
			for j := i + 1; j < len(groupedPositions); j++ {
				node1 := groupedPositions[i]
				node2 := groupedPositions[j]

				antinode1 := findAntinode(node1, node2)
				antinode2 := findAntinode(node2, node1)

				// fmt.Printf("node1: %d, node2 %d, antinode %d \n", node1, node2, antinode1)
				// fmt.Printf("node2: %d, node1 %d, antinode %d \n", node2, node1, antinode2)
				if isNodeInsideBoard(board, antinode1) {
					createdAntinodes[antinode1] = true
				}

				if isNodeInsideBoard(board, antinode2) {
					createdAntinodes[antinode2] = true
				}
			}
		}
	}

	// fmt.Println(createdAntinodes)
	return len(createdAntinodes)
}

func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day8/input1")
	check(err)

	board := strings.Split(string(fileData), "\n")
	groupedAntennas := groupAllAntennas(board)
	createdAntinodes := make(map[[2]int]bool)
	// fmt.Println(groupedAntennas)

	for _, groupedPositions := range groupedAntennas {
		for i := 0; i < len(groupedPositions); i++ {
			for j := i + 1; j < len(groupedPositions); j++ {
				node1 := groupedPositions[i]
				node2 := groupedPositions[j]

				nodesDifference := findNodesDifference(node1, node2)
				nextNode := findNextNode(node1, nodesDifference)

				for {
					if !isNodeInsideBoard(board, nextNode) {
						break
					}

					createdAntinodes[nextNode] = true
					nextNode = findNextNode(nextNode, nodesDifference)
				}

				nodesDifference = findNodesDifference(node2, node1)
				nextNode = findNextNode(node2, nodesDifference)

				for {
					if !isNodeInsideBoard(board, nextNode) {
						break
					}

					createdAntinodes[nextNode] = true
					nextNode = findNextNode(nextNode, nodesDifference)
				}
			}
		}
	}

	// fmt.Println(createdAntinodes)
	return len(createdAntinodes)
}

func groupAllAntennas(board []string) map[byte][][2]int {
	groupedNodes := make(map[byte][][2]int)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			antennaType := board[i][j]
			if groupedNodes[antennaType] == nil {
				groupedNodes[antennaType] = make([][2]int, 0)
			}

			groupedNodes[antennaType] = append(groupedNodes[antennaType], [2]int{i, j})
		}
	}
	return groupedNodes
}

// {3, 4} , {5, 5} -> {7, 6}
// {5, 5}, {3, 4} -> {1, 3}
func findAntinode(node1, node2 [2]int) [2]int {
	dx := node2[1] - node1[1]
	dy := node2[0] - node1[0]

	return [2]int{node2[0] + dy, node2[1] + dx}
}

func findNodesDifference(node1, node2 [2]int) [2]int {
	return [2]int{node2[0] - node1[0], node2[1] - node1[1]}
}

func findNextNode(node, difference [2]int) [2]int {
	return [2]int{node[0] + difference[0], node[1] + difference[1]}
}

func isNodeInsideBoard(board []string, node [2]int) bool {
	if node[0] < 0 || node[1] < 0 || node[0] >= len(board) || node[1] >= len(board[1]) {
		return false
	}

	return true
}
