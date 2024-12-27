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


	// P1(&input, 25)
	P2(&input, 75)
}

func P2(input *[]string, blinks int) {
	total := 0
	input_map := make(map[string]int) 
	// initialize map
	for _, c := range(*input) {
		input_map = add_to_map(input_map, string(c), 1)
	}
	for i:= 0; i < blinks; i++ {
		// fmt.Println(i)
		input_map = blink_map(&input_map)
	}
	for _, v := range(input_map) {
		total += v
	}
	fmt.Println(total, blinks)
}

func blink_map(input_map *map[string]int) map[string]int {
	output_map := make(map[string]int) 
	for k := range *input_map {
		char := strings.TrimLeft(strings.TrimSpace(k), "0")
		count := (*input_map)[k]
		if count == 0 {
			delete(*input_map, k)
		} else {
				if char == "" {
					output_map = add_to_map(output_map, "1", count)
				} else if len(char) %2 == 0 {
					mid := len(char)/2
					left := strings.TrimLeft(strings.TrimSpace(char[mid:]), "0")
					right := strings.TrimLeft(strings.TrimSpace(char[:mid]), "0")
					if left == "" {
						left = "0"
					}
					if right == "" {
						right = "0"
					}
					output_map = add_to_map(output_map, left, count)
					output_map = add_to_map(output_map, right, count)
				} else  {
					cInt, err := strconv.Atoi(char)
					if err != nil {
						panic(err)
					}
					output_map = add_to_map(output_map, strconv.Itoa(cInt*2024), count)
				}
			
		}
	}
	return output_map
}


func add_to_map(input map[string]int, key string, count int) map[string]int {
	_, exists := input[key]
	if exists {
		input[key] += count
	} else {
		input[key] = count
	}
	return input
}

func blink2(char string, blinks int) []string {
	if blinks <= 0 {
		return []string{char}
	}
	char = strings.TrimLeft(strings.TrimSpace(char), "0")
	if char == "" {
		return blink2("1", blinks-1)
	} else if len(char) %2 == 0 {
		mid := len(char)/2
		return append(blink2(char[mid:], blinks-1), blink2(char[:mid], blinks-1)...)
	} else  {
		cInt, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		return blink2(strconv.Itoa(cInt*2024), blinks-1)
	}
}

func P1(input *[]string, blinks int) {
	// fmt.Println(*input)
	newLine := *input
	for i:= 0; i < blinks; i++ {
		newLine = blink(&newLine)
		// fmt.Println(i, len(newLine))
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