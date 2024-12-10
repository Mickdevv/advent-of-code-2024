package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var trailHeads [][2]int
var trail [][2]int

func main() {
	file, err := os.Open("Day_10/input_test.txt")
	// file, err := os.Open("Day_10/input.txt")
	if err != nil {
		panic(err)
	}

	var line string
	var input [][]int
	var cInt int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		lineInt := []int{}
		for _, c := range(line) {
			cInt, _ = strconv.Atoi(string(c))
			lineInt = append(lineInt, cInt)
		}
		input = append(input, lineInt)
	}

	P1(&input)
	// P2(&input)
}

func P2(input *[][]int) {
	fmt.Println(*input)
}

func P1(input *[][]int) {
	for lineIndex, line := range(*input) {
		for elemIndex, elem := range(line) {
			if elem == 0 {
				fmt.Println(findTrailheads(input, [2]int{elemIndex, lineIndex}))
			}
		}
	}
	fmt.Println(trailHeads)
}

func findTrailheads(input *[][]int, sPos [2]int) [2]int {
	var trail [][2]int
	currentPos := sPos
	currentValue := getItemAtIndex(input, sPos)
	possibleMoves := getMoveMatrix(currentPos)
	moves := 0

	for currentValue < 9 && moves < 100 {
		for _, move := range(possibleMoves) {
			if isValidMove(input, currentValue, move, &trail) {
				currentPos = moveFunc(input, move)
			}
		}
	}
	// fmt.Println("moves: ", moves, "Value : ", getItemAtIndex(input, currentPos), "currentPos: ", currentPos)
	return currentPos
}

func isValidMove(input *[][]int, sInt int, mPos [2]int, trail *[][2]int) bool {
	if mPos[1] < len(*input) && mPos[1] >= 0 && mPos[0] < len((*input)[0]) && mPos[0] >= 0 && (sInt-getItemAtIndex(input, mPos) == 1 || sInt-getItemAtIndex(input, mPos) == -1) && !isPosInTrail(trail, mPos) {
		return true
	}
	return false
}

func getItemAtIndex(input *[][]int, pos [2]int) int {
	if pos[1] < 0 || pos[0] < 0 || pos[1] >= len(*input) || pos[0] >= len((*input)[0]) {
		return -2
	}
	return (*input)[pos[1]][pos[0]]
}

func isPosInTrail(trail *[][2]int, pos [2]int) bool {
	for i := len(*trail)-1; i > 0; i-- {
		if pos == (*trail)[i] {
			return true
		}
	}
	return false
}

func getMoveMatrix(sPos [2]int) [4][2]int {
	upPos := [2]int{sPos[0], sPos[1]+1}
	downPos := [2]int{sPos[0], sPos[1]-1}
	leftPos := [2]int{sPos[0]-1, sPos[1]}
	rightPos := [2]int{sPos[0]+1, sPos[1]}
	return [4][2]int{upPos, downPos, leftPos, rightPos}
}

func moveFunc(input *[][]int, currentPos [2]int) [2]int {
	count := 0
	for _, move := range(getMoveMatrix(currentPos)) {
		count = 0
		fmt.Println(1)
		if isValidMove(input, getItemAtIndex(input, move), move, &trail) {
			count ++
			fmt.Println(2)
			if getItemAtIndex(input, move) != 9 && isValidMove(input, getItemAtIndex(input, move), move, &trail) {
				trail = append(trail, move)
				fmt.Println(trail)
				return moveFunc(input, move)
			} else if isValidMove(input, getItemAtIndex(input, move), move, &trail) && getItemAtIndex(input, move) == 9 {
				fmt.Println("Found ! ", move, getItemAtIndex(input, move))
				trailHeads = append(trailHeads, move)
				return move
			}
		} else {
			fmt.Println("Dead end !")
		}
		if count == 0 {
			break
		}
	}
	return [2]int{-1, -1}
}