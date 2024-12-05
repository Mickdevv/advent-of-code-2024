package main

import (
	"bufio"
	"fmt"
	"os"
)

var notFound = [4][2]int{{0,0},{0,0},{0,0},{0,0}}

func main(){
	file, err := os.Open("Day_4/input.txt")
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
	P2(table)
}

func P2(table []string) {
	total := 0
	for lineIndex, lineValue := range(table) {

		for charIndex, charValue := range(lineValue) {
			if string(charValue) == "A" && charIndex > 0 && charIndex < len(lineValue) -1  && lineIndex > 0 && lineIndex < len(table) -1 {
				fmt.Println("==", charIndex, lineIndex)
				total += searchAroundMas(table, charIndex, lineIndex)
			}
		}
	}
	fmt.Println(total)
}

func searchAroundMas(table []string, x int, y int) int {
	total := 0

	if string(table[y+1][x+1]) == "M" && string(table[y+1][x-1]) == "M" && string(table[y-1][x-1]) == "S" && string(table[y-1][x+1]) == "S" {
		fmt.Println(x, y)
		total ++
	} else if string(table[y+1][x+1]) == "M" && string(table[y+1][x-1]) == "S" && string(table[y-1][x-1]) == "S" && string(table[y-1][x+1]) == "M" {
		fmt.Println(x, y)
		total ++
	} else if string(table[y+1][x+1]) == "S" && string(table[y+1][x-1]) == "S" && string(table[y-1][x-1]) == "M" && string(table[y-1][x+1]) == "M" {
		fmt.Println(x, y)
		total ++
	} else if string(table[y+1][x+1]) == "S" && string(table[y+1][x-1]) == "M" && string(table[y-1][x-1]) == "M" && string(table[y-1][x+1]) == "S" {
		fmt.Println(x, y)
		total ++
	} 
	return total

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
	// fmt.Println(x)
	// fmt.Println(table, x, y)
	total := 0

	//check inline to the right
	if x + len("XMAS") <= len(table[y]) && table[y][x:x+len("XMAS")] == "XMAS" {
		total ++
		// fmt.Println(table[y], "right", x, y)
	}
	//check inline to the left
	if x+1 - len("XMAS") >= 0 && table[y][x+1 - len("XMAS"):x+1] == "SAMX" {
		total ++
		// fmt.Println(table[y], "left", x, y)
	}


	XMAS := "XMAS"
	// SAMX := "SAMX"

	upRight := ""
	upLeft := ""
	downRight := ""
	downLeft := ""
	up := ""
	down := ""

	for i := 0; i < len(XMAS); i++ {
		if y >= i {
			up += string(table[y - i][x])
		}
		if y + i < len(table) {
			down += string(table[y + i][x])
		}
		if y >= i && x + i < len(table[y]) {
			upRight += string(table[y - i][x + i])
		}
		if y >= i && x >= i {
			upLeft += string(table[y - i][x - i])
		}
		if y + i < len(table) && x + i < len(table[y]) {
			downRight += string(table[y + i][x + i])
		}
		if y + i < len(table) && x >= i {
			downLeft += string(table[y + i][x - i])
		}
	}

	if up == XMAS {
		// fmt.Println("up", x, y, up)
		total++
	}
	if down == XMAS {
		// fmt.Println("down", x, y, down)
		total++
	}
	if upRight == XMAS {
		// fmt.Println("upRight", x, y, upRight)
		total++
	}
	if upLeft == XMAS {
		// fmt.Println("upLeft", x, y, upLeft)
		total++
	}
	if downRight == XMAS {
		// fmt.Println("downRight", x, y, downRight)
		total++
	}
	if downLeft == XMAS {
		// fmt.Println("downLeft", x, y, downLeft)
		total++
	}
	return total
}