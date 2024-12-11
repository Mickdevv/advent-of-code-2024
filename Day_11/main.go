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
	// file, err := os.Open("Day_11/input.txt")
	file, err := os.Open("Day_11/input.txt")
	if err != nil {
		panic(err)
	}

	var input []string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input = strings.Split(scanner.Text(), " ")


	// P1(&input, 75)
	P2(&input, 6)
}

func P2(input *[]string, blinks int) {
	total := 0
	for _, c := range(*input) {
		newLine := []string{string(c)}
		total += len(blink2(&newLine, blinks))
	}
	fmt.Println(total)
}

func P1(input *[]string, blinks int) {
	// fmt.Println(*input)
	newLine := *input
	for i:= 0; i < blinks; i++ {
		newLine = blink(&newLine)
		fmt.Println(i, len(newLine))
	}
	fmt.Println(len(newLine))
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

func blink2(line *[]string, blinks int) []string {
	var newLine []string
	for _, c := range(*line) {
		char := strings.TrimSpace(string(c))
		charInt, _ := strconv.Atoi(char)
		char = strconv.Itoa(charInt)
		if char == "0" {
			newLine = append(newLine, "1")
		} else if len(char) %2 ==0 {
			newLine = append(newLine, char[len(char)/2:])
			newLine = append(newLine, char[:len(char)/2])
			// fmt.Println("|", char[len(char)/2:],"|", char[:len(char)/2],"|")
		} else  {
			cInt, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			newLine = append(newLine, strconv.Itoa(cInt*2024))
		}
	
	}
	if blinks > 0 {
		return blink2(&newLine, blinks-1)
	} else {
		return newLine
	}
}