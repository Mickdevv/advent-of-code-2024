package main

import (
	"bufio"
	"fmt"
	"os"
)

var notFound = [4][2]int{{0,0},{0,0},{0,0},{0,0}}

func main(){
	file, err := os.Open("Day_4/input_test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var table []string
	for scanner.Scan() {
		line := scanner.Text()
		table = append(table, line)
	}

	P1(table)
}

func P1(table []string) {
	total := 0
	for lineIndex, lineValue := range(table) {
		for charIndex, charValue := range(lineValue) {
			if string(charValue) == "X" {
				total += searchAround(table, charIndex, lineIndex)
			}
		}
	}
	fmt.Println(total)
}

func searchAround(table []string, x int, y int) int {
	// fmt.Println(table, x, y)
	total := 0

	//check inline to the right
	if x + len("XMAS") < len(table[y]) && table[y][x:x+len("XMAS")] == "XMAS" {
		total ++
		fmt.Println(table[y], "XMAS")
	}
	//check inline to the left
	if x - len("XMAS") >= 0 && table[y][x - len("XMAS")+1:x+1] == "SAMX" {
		total ++
		fmt.Println(table[y], "SAMX")
	}

	XMAS := "XMAS"

	upRight := 0
	upLeft := 0
	downRight := 0
	downLeft := 0

	for i := 0; i < len(XMAS); i++ {
		if i < len(table) && i < len(table[y]) && table[y][x] == XMAS[i] {
			upRight ++
		}
		if i < len(table) && i < len(table[y]) && table[y][x] == XMAS[i] {
			upLeft ++
		}
		if i < len(table) && i < len(table[y]) && table[y][x] == XMAS[i] {
			downRight ++
		}
		if i < len(table) && i < len(table[y]) && table[y][x] == XMAS[i] {
			downLeft ++
		}
	}

	if upRight == 4 {
		total++
	}
	if upLeft == 4 {
		total++
	}
	if downRight == 4 {
		total++
	}
	if downLeft == 4 {
		total++
	}
	return total

}