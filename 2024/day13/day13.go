package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Game struct {
	ax int
	ay int
	bx int
	by int
	x  int
	y  int
}

const MAX_USES_OF_BUTTON = 100
const COORDS_INCREASED = 10000000000000

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
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day13/input1")
	check(err)

	games := strings.Split(string(fileData), "\n\n")
	res := 0

	for _, gameString := range games {
		game := parseGame(gameString)
		points, success := gameResult(game)

		if success {
			res += points
		}
	}

	return res
}

func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day13/input1")
	check(err)

	games := strings.Split(string(fileData), "\n\n")
	res := 0

	for _, gameString := range games {
		game := parseGame(gameString)
		game.x += COORDS_INCREASED
		game.y += COORDS_INCREASED

		points, success := gameResult(game)

		if success {
			res += points
		}
	}

	return res
}

func parseGame(game string) Game {
	resGame := Game{}
	gameRows := strings.Split(game, "\n")

	buttonsRegexp, err := regexp.Compile(`X\+(\d+), Y\+(\d+)`)
	check(err)

	prizeRegexp, err := regexp.Compile(`X\=(\d+), Y\=(\d+)`)
	check(err)

	firstButton := buttonsRegexp.FindStringSubmatch(gameRows[0])
	secondButton := buttonsRegexp.FindStringSubmatch(gameRows[1])
	prize := prizeRegexp.FindStringSubmatch(gameRows[2])

	resGame.ax = parseNumber(firstButton[1])
	resGame.ay = parseNumber(firstButton[2])
	resGame.bx = parseNumber(secondButton[1])
	resGame.by = parseNumber(secondButton[2])
	resGame.x = parseNumber(prize[1])
	resGame.y = parseNumber(prize[2])

	return resGame
}

func parseNumber(number string) int {
	res, err := strconv.Atoi(number)
	check(err)

	return res
}

func gameResult(game Game) (int, bool) {
	A := mat.NewDense(
		2, 2, []float64{float64(game.bx), float64(game.ax), float64(game.by), float64(game.ay)},
	)
	b := mat.NewVecDense(2, []float64{float64(game.x), float64(game.y)})

	var x mat.VecDense
	if err := x.SolveVec(A, b); err != nil {
		return 0, false
	}

	intI := int(math.Round(x.RawVector().Data[0]))
	intJ := int(math.Round(x.RawVector().Data[1]))
	if (game.bx*intI)+(game.ax*intJ) != game.x ||
		(game.by*intI)+(game.ay*intJ) != game.y {

		return 0, false
	}

	return intI + (intJ * 3), true
}
