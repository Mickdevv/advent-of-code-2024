package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type robot struct {
	p [2]int;
	v [2]int;
}

func main() {
	// file, err := os.Open("Day_14/input.txt")
	file, err := os.Open("Day_14/input_test.txt")
	if err != nil {
		panic(err)
	}

	var robots []robot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		
		pString := strings.Split(strings.Split(strings.Split(line, " ")[0], "=")[1], ",")
		vString := strings.Split(strings.Split(strings.Split(line, " ")[1], "=")[1], ",")

		r := robot{p: [2]int(strArrToIntArr(pString)), v: [2]int(strArrToIntArr(vString))}
		robots = append(robots, r)
		
	}
	fmt.Println(robots)
	for a := range(1) {
		a++
		P1(&robots, 11, 7, 100)
	}
}

func P1(robots *[]robot, boardX int, boardY int, time int) {
	// fmt.Println(robots)
	for i, robot := range(*robots) {
		fmt.Println("Before moves: ", robot, robot.p[0] + robot.v[0] * time, robot.p[1] + robot.v[1] * time)
		
		robot.p[0] += robot.v[0] * time
		robot.p[1] += robot.v[1] * time
		fmt.Println("After moves: ", robot)

		if robot.p[0] < 0 {
			robot.p[0] = boardX - int(math.Abs(float64(robot.p[0]))) % boardX-1
		} else {
			robot.p[0] = robot.p[0] % boardX
		}

		if robot.p[1] < 0 {
			robot.p[1] = boardY - int(math.Abs(float64(robot.p[1]))) % boardY-1
		} else {
			robot.p[1] = robot.p[1] % boardY
		}
		(*robots)[i] = robot
		fmt.Println(robot.p)
		fmt.Println("-----")
	}

	middleX := (boardX-1)/2
	middleY := (boardY-1)/2

	var quadrants [4]int

	fmt.Println(middleX, middleY)

	for _, robot := range(*robots) {

		if robot.p[0] < middleX && robot.p[1] < middleY {
			//top left
			quadrants[0] ++
			fmt.Println(robot.p)
		} else if robot.p[0] > middleX && robot.p[1] < middleY {
			//top right
			quadrants[1] ++
		} else if robot.p[0] < middleX && robot.p[1] > middleY {
			//bottom left
			quadrants[2] ++
		} else if robot.p[0] > middleX && robot.p[1] > middleY {
			//bottom right
			quadrants[3] ++
		} 
	}
	total := 1
	for _, q := range(quadrants) {
		if q > 0 {
			total *= q
		}
	}
	fmt.Println(quadrants, total)
	// fmt.Println(len(*robots))
}

func strArrToIntArr(strArr []string) []int {
	var intArr []int
	for _, v := range(strArr) {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		intArr = append(intArr, i)
	}
	return intArr
}