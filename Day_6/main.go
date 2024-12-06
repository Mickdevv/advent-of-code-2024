package main

import (
	"bufio"
	"fmt"
	"os"
)

type move struct {
    location [2]int
    direction string
}

func main() {
	var input []string

	file, err := os.Open("Day_6/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	// fmt.Println(input)

	P1(&input)
	P2(&input)
}

func P2(input *[]string) {
	total := 0
	guardCurrentPos, guard := guardStartingPos(*input)
	
	for lineindex, line := range(*input) {
		for i, c := range(line) {
			if !isObstacleOrGuard(string(c)) {
				tempTable := append([]string{}, *input...)
				s := []byte(tempTable[lineindex])
				s[i] = '#'
				tempTable[lineindex] = string(s)
				// fmt.Println(tempTable, guardCurrentPos, guard, total)
				total += detectLoop(&tempTable, guardCurrentPos, guard)

			}
		}
	}
	fmt.Println(total)
}

func detectLoop(input *[]string, guardCurrentPos [2]int, guard string) int {
	guardStartingPos := guardCurrentPos
	var moves []move
	var move move
	var path [][2]int
	// fmt.Println(guardStartingPos(input))

	for guardCurrentPos[0] < len(*input)-1 && guardCurrentPos[1] < len((*input)[0])-1 && guardCurrentPos[1] > 0 &&guardCurrentPos[0] > 0 && !moveInPath(moves, move) {
		if move.location != guardStartingPos {
			moves = append(moves, move)
		}
		
		if guard == ">" {
			guardCurrentPos, guard, _, move = detectMoveRight(input, guardCurrentPos, guard)
		} else if guard == "V" {
			guardCurrentPos, guard, _, move = detectMoveDown(input, guardCurrentPos, guard)
		} else if guard == "<" {
			guardCurrentPos, guard, _, move = detectMoveLeft(input, guardCurrentPos, guard)
		} else if guard == "^" {
			guardCurrentPos, guard, _, move = detectMoveUp(input, guardCurrentPos, guard)
		}
		if !stepInPath(path, move.location) {
			path = append(path, move.location)
		}
		// fmt.Println(move, moves, moveInPath(moves, move))

	}
	// fmt.Println(len(path) +1, moves, move)
	if moveInPath(moves, move) {
		return 1
	}
	return 0
}

func P1(input *[]string) {
	count := 0
	var addToCount int
	var path [][2]int
	var step [2]int
	guardCurrentPos, guard := guardStartingPos(*input)
	// fmt.Println(guardStartingPos(input))

	for guardCurrentPos[0] < len(*input)-1 && guardCurrentPos[1] < len((*input)[0])-1 && guardCurrentPos[1] > 0 &&guardCurrentPos[0] > 0   {
		addToCount = 0
		if guard == ">" {
			guardCurrentPos, guard, addToCount, step = detectRight(*input, guardCurrentPos, guard)
			// fmt.Println(guardCurrentPos, guard)
		} else if guard == "V" {
			guardCurrentPos, guard, addToCount, step = detectDown(*input, guardCurrentPos, guard)
		} else if guard == "<" {
			guardCurrentPos, guard, addToCount, step = detectLeft(*input, guardCurrentPos, guard)
		} else if guard == "^" {
			guardCurrentPos, guard, addToCount, step = detectUp(*input, guardCurrentPos, guard)
		}
		if !stepInPath(path, step) {
			path = append(path, step)
			count += addToCount
		}
	}
	fmt.Println(len(path)+1)
}

func guardStartingPos(input []string) ([2]int, string) {
	for lineindex, line := range(input) {
		for i, c := range(line) {
			if string(c) == ">" || string(c) == "^" || string(c) == "<" || string(c) == "V" {
				return [2]int{i, lineindex}, string(c)
			}
		}
	}
	return [2]int{0, 0}, ""
}

func detectMoveRight(input *[]string, guardCurrentPos [2]int, guard string) ([2]int, string, int, move) {

	if string((*input)[guardCurrentPos[1]][guardCurrentPos[0] +1]) == "#" {
		// fmt.Println("obstacle detected to the right")
		move := move{location: guardCurrentPos, direction: "V"}
		return guardCurrentPos, "V", 0, move
	} else {
		guardCurrentPos[0] += 1
		move := move{location: guardCurrentPos, direction:guard}
		return guardCurrentPos, guard, 1, move
	}
}

func detectMoveLeft(input *[]string, guardCurrentPos [2]int, guard string) ([2]int, string, int, move) {
	if string((*input)[guardCurrentPos[1]][guardCurrentPos[0] -1]) == "#" {
		// fmt.Println("obstacle detected to the left")
		move := move{location: guardCurrentPos, direction: "^"}
		return guardCurrentPos, "^", 0, move
	} else {
		guardCurrentPos[0] -= 1
		move := move{location: guardCurrentPos, direction:guard}
		return guardCurrentPos, guard, 1, move
	}
}

func detectMoveUp(input *[]string, guardCurrentPos [2]int, guard string) ([2]int, string, int, move) {
	if string((*input)[guardCurrentPos[1]-1][guardCurrentPos[0]]) == "#" {
		// fmt.Println("obstacle detected above")
		move := move{location: guardCurrentPos, direction: ">"}
		return guardCurrentPos, ">", 0, move
	} else {
		guardCurrentPos[1] -= 1
		move := move{location: guardCurrentPos, direction:guard}
		return guardCurrentPos, guard, 1, move
	}
}


func detectMoveDown(input *[]string, guardCurrentPos [2]int, guard string) ([2]int, string, int, move) {
	if string((*input)[guardCurrentPos[1]+1][guardCurrentPos[0]]) == "#" {
		// fmt.Println("obstacle detected below")
		move := move{location: guardCurrentPos, direction: "<"}
		return guardCurrentPos, "<", 0, move
	} else {
		guardCurrentPos[1] += 1
		move := move{location: guardCurrentPos, direction:guard}
		return guardCurrentPos, guard, 1, move
	}
}

func detectRight(input []string, guardCurrentPos [2]int, guard string) ([2]int, string, int, [2]int) {
	guardOldPos := guardCurrentPos
	if string(input[guardCurrentPos[1]][guardCurrentPos[0] +1]) == "#" {
		// fmt.Println("obstacle detected to the right")
		return guardCurrentPos, "V", 0, guardOldPos
	} else {
		guardCurrentPos[0] += 1
		return guardCurrentPos, guard, 1, guardOldPos
	}
}

func detectLeft(input []string, guardCurrentPos [2]int, guard string) ([2]int, string, int, [2]int) {
	guardOldPos := guardCurrentPos
	if string(input[guardCurrentPos[1]][guardCurrentPos[0] -1]) == "#" {
		// fmt.Println("obstacle detected to the left")
		return guardCurrentPos, "^", 0, guardOldPos
	} else {
		guardCurrentPos[0] -= 1
		return guardCurrentPos, guard, 1, guardOldPos
	}
}

func detectUp(input []string, guardCurrentPos [2]int, guard string) ([2]int, string, int, [2]int) {
	guardOldPos := guardCurrentPos
	if string(input[guardCurrentPos[1]-1][guardCurrentPos[0]]) == "#" {
		// fmt.Println("obstacle detected above")
		return guardCurrentPos, ">", 0, guardOldPos
	} else {
		guardCurrentPos[1] -= 1
		return guardCurrentPos, guard, 1, guardOldPos
	}
}


func detectDown(input []string, guardCurrentPos [2]int, guard string) ([2]int, string, int, [2]int) {
	guardOldPos := guardCurrentPos
	if string(input[guardCurrentPos[1]+1][guardCurrentPos[0]]) == "#" {
		// fmt.Println("obstacle detected below")
		return guardCurrentPos, "<", 0, guardOldPos
	} else {
		guardCurrentPos[1] += 1
		return guardCurrentPos, guard, 1, guardOldPos
	}
}

func stepInPath(path [][2]int, step [2]int) bool {
	for _, s := range(path) {
		if step == s {
			return true
		}
	}
	return false
}

func moveInPath(moves []move, move move) bool {
	for _, m := range(moves) {
		if move == m {
			return true
		}
	}
	return false
}

func isObstacleOrGuard(c string) bool {
	if c == "V" ||c == ">" ||c == "<" ||c == "^" ||c == "#" {
		return true
	}
	return false
}