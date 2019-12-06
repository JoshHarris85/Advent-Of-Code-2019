package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func readIntCode() ([]string, []string) {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	text := string(content)
	// Move strings to array
	moves := strings.Split(text, "\n")
	x := strings.Split(moves[0], ",")
	y := strings.Split(moves[1], ",")

	return x, y
}

func getMoveSet(moves []string) [][3]int {
	var moveSet [][3]int
	lastMove := [3]int{0, 0, 0}
	steps := 0

	for i := 0; i < len(moves); i++ {
		move := moves[i][:1]
		count, _ := strconv.Atoi(moves[i][1:])

		switch move {
		case "U":
			for n := 0; n < count; n++ {
				currentMove := [3]int{lastMove[0], lastMove[1] + 1, steps + 1}
				moveSet = append(moveSet, currentMove)
				lastMove = currentMove
				steps++
			}
		case "D":
			for n := 0; n < count; n++ {
				currentMove := [3]int{lastMove[0], lastMove[1] - 1, steps + 1}
				moveSet = append(moveSet, currentMove)
				lastMove = currentMove
				steps++
			}
		case "L":
			for n := 0; n < count; n++ {
				currentMove := [3]int{lastMove[0] - 1, lastMove[1], steps + 1}
				moveSet = append(moveSet, currentMove)
				lastMove = currentMove
				steps++
			}
		case "R":
			for n := 0; n < count; n++ {
				currentMove := [3]int{lastMove[0] + 1, lastMove[1], steps + 1}
				moveSet = append(moveSet, currentMove)
				lastMove = currentMove
				steps++
			}
		}
	}

	return moveSet
}

func calculateMoves() [][3]int {
	movesOne, movesTwo := readIntCode()

	moveSetOne := getMoveSet(movesOne)
	moveSetTwo := getMoveSet(movesTwo)

	var overlappingMoves [][3]int

	for x := 0; x < len(moveSetOne); x++ {
		for y := 0; y < len(moveSetTwo); y++ {
			if moveSetOne[x][0] == moveSetTwo[y][0] && moveSetOne[x][1] == moveSetTwo[y][1] {
				currentMove := [3]int{moveSetOne[x][0], moveSetOne[x][1], moveSetOne[x][2] + moveSetOne[y][2]}
				overlappingMoves = append(overlappingMoves, currentMove)
			}
		}
	}

	return overlappingMoves
}

func partOne() float64 {
	moves := calculateMoves()

	var closest float64
	closest = math.Abs(float64(moves[0][0])) + math.Abs(float64(moves[0][1]))

	for x := 0; x < len(moves); x++ {
		distance := math.Abs(float64(moves[x][0])) + math.Abs(float64(moves[x][1]))

		if distance < closest {
			closest = distance
		}
	}

	return closest
}

func partTwo() int {
	moves := calculateMoves()
	closest := moves[0][2]

	for x := 0; x < len(moves); x++ {
		if moves[x][2] < closest {
			closest = moves[x][2]
		}
	}

	return closest
}

func main() {
	fmt.Println("The answer to part 1 is:", partOne())
	fmt.Println("The answer to part 2 is:", partTwo())
}
