package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var trailHeads [][2]int
var trail [][2]int

func main() {
	file, err := os.Open("Day_11/input.txt")
	// file, err := os.Open("Day_11/input_test.txt")
	if err != nil {
		panic(err)
	}

	var input []string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input = strings.Split(scanner.Text(), " ")


	// P1(&input, 75)
	P2(&input, 25)
}

func P2(input *[]string, blinks int) {
	total := 0
	for _, c := range(*input) {
		char := string(c)
		processed_line := blink2(&char, blinks)
		total += len(processed_line)
	}
	fmt.Println(total, blinks)
}

func blink2(char *string, blinks int) []string {
	if blinks <= 0 {
		return []string{*char}
	}
	newChar := strings.TrimLeft(strings.TrimSpace(*char), "0")
	if newChar == "" {
		newChar = "1"
		return blink2(&newChar, blinks-1)
	} else if len(newChar) %2 == 0 {
		left := newChar[len(newChar)/2:]
		right := newChar[:len(newChar)/2]
		return append(blink2(&left, blinks-1), blink2(&right, blinks-1)...)
	} else  {
		cInt, err := strconv.Atoi(newChar)
		if err != nil {
			panic(err)
		}
		newChar = strconv.Itoa(cInt*2024)
		return blink2(&newChar, blinks-1)
	}
}

func P1(input *[]string, blinks int) {
	// fmt.Println(*input)
	newLine := *input
	for i:= 0; i < blinks; i++ {
		newLine = blink(&newLine)
		// fmt.Println(i, len(newLine))
	}
	// fmt.Println(len(newLine))
}

func blink(line *[]string) []string {
	var newLine []string
	for _, c := range(*line) {
		char := strings.TrimLeft(strings.TrimSpace(string(c)), "0")
		if char == "" {
			newLine = append(newLine, "1")
		} else if len(char) %2 ==0 {
			newLine = append(newLine, strings.TrimSpace(char[len(char)/2:]))
			newLine = append(newLine, strings.TrimSpace(char[:len(char)/2]))
		} else  {
			cInt, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			newLine = append(newLine, strconv.Itoa(cInt*2024))
		}
	}
	return newLine
}