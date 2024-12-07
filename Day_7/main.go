package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var availableOperators []string = []string{"*", "/"}

func main() {
	file, err := os.Open("Day_7/input_test.txt")
	if err != nil {
		panic(err)
	}
	
	scanner := bufio.NewScanner(file)
	var input [][]int

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ":")
		var lineFinal []int
		answer, err := strconv.Atoi(lineSplit[0])
		if err != nil {
			panic(err)
		}
		lineFinal = append(lineFinal, answer)
		for _, v := range(strings.Split(lineSplit[1], " ")) {
			answer, err := strconv.Atoi(v)
			if err != nil {
				// fmt.Println("Error converting ", v)
			}
			lineFinal = append(lineFinal, answer)
		}
		input = append(input, lineFinal)
	}

	P1(&input)
}
func P1(input *[][]int) {
	fmt.Println(*input)
	total := 0

	for _, line := range(*input) {
		total += checkLine(line)
	}
	fmt.Println(total)
}

func checkLine(line []int) int {
	fmt.Println(line)
	// answer := line[0]
	// var checkAnswer int
	var operators []string
	for range(len(line)-2) {
		
	}
	fmt.Println(operators)
	return 0
}