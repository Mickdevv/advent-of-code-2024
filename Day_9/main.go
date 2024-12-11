package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
	id int
    block_size int
    block_free_space int
}

func main() {
	// file, err := os.Open("Day_9/input_test.txt")
	file, err := os.Open("Day_9/input.txt")
	if err != nil {
		panic(err)
	}
	
	var input []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	input_array_string := strings.Split(input[0], "")
	var input_array_int []int
	
	for _, c := range(input_array_string) {
		cInt, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println(err, string(c))
		}
		input_array_int = append(input_array_int, cInt)
	}

	P1(&input_array_int)
	P2(&input_array_int)
}

func P1(input *[]int) {
	var files []file
	for i := 0; i < len(*input); i+=2 {
		if i+1 >= len(*input) {
			files = append(files, file{id: i/2, block_size: (*input)[i], block_free_space: 0})
			} else {
			files = append(files, file{id: i/2, block_size: (*input)[i], block_free_space: (*input)[i+1]})
		}
	}
	// fmt.Println(files)
	
	checkSum := file_list_to_file_string_1(&files)
	fmt.Println(checkSum)
}

func P2(input *[]int) {
	var files []file
	for i := 0; i < len(*input); i+=2 {
		if i+1 >= len(*input) {
			files = append(files, file{id: i/2, block_size: (*input)[i], block_free_space: 0})
			} else {
			files = append(files, file{id: i/2, block_size: (*input)[i], block_free_space: (*input)[i+1]})
		}
	}
	// fmt.Println(files)
	
	checkSum := file_list_to_file_string_2(&files)
	fmt.Println(checkSum)
}

func file_list_to_file_string_2(files *[]file) int {
	var answer [][]int

	for _, file := range(*files) {
		// fmt.Println(file.block_free_space, file.block_size, file.id)
		var block_full []int
		for i := 0; i < file.block_size; i++ {
			block_full = append(block_full, file.id)
		}
		var block_empty []int
		for i := 0; i < file.block_free_space; i++ {
			block_empty = append(block_empty, -1)
		}
		if len(block_full) > 0 {
			answer = append(answer, block_full)
		}
		if len(block_empty) > 0 {
			answer = append(answer, block_empty)
		}
	}

	// fmt.Println(answer)

	leftPos := 0
	var currentId int
	for i := len(answer)-1; i > 0; i-- {
		if answer[i][0] != -1 {
			currentId = answer[i][0]
			break
		}
		// fmt.Println(answer, i, answer[i][0], currentId)
	}
	// fmt.Println(answer, currentId)
	
	for currentId != 0 {
		rightPos := len(answer)-1
		for answer[rightPos][0] != currentId {
			rightPos --
		}
		for !(answer[leftPos][0] == -1 && len(answer[leftPos]) >= len(answer[rightPos])) && leftPos < len(answer)-1 {
			leftPos ++
		} 
		// fmt.Println()
		if answer[leftPos][0] == -1 && answer[rightPos][0] == currentId && len(answer[leftPos]) >= len(answer[rightPos]) && leftPos < rightPos {
			// fmt.Println("==========")
			// fmt.Println(answer, answer[leftPos], answer[rightPos], currentId)
			rightPosLen := len(answer[rightPos])
			for i := 0; i < rightPosLen; i++ {
				answer[leftPos][i], answer[rightPos][i] = answer[rightPos][i], answer[leftPos][i]
			}
			// fmt.Println(answer, answer[leftPos], answer[rightPos])
			if len(answer[leftPos]) != len(answer[rightPos]) {
				answer = reGroup(answer)
			}
			// fmt.Println(answer)
			// fmt.Println("==========")
			// answer[leftPos], answer[rightPos] = answer[rightPos], answer[leftPos]
		}
		leftPos = 0
		currentId --
	}

	// fmt.Println(answer)

	checkSum := 0
	var flatAnswer []int
	for _, a := range(answer) {
		flatAnswer = append(flatAnswer, a...)
	}

	for i, c := range(flatAnswer) {
		if c != -1 {
			checkSum += c*i
		}
	}

	return checkSum
}

func file_list_to_file_string_1(files *[]file) int {
	var answer []int

	for _, file := range(*files) {
		// fmt.Println(file.block_free_space, file.block_size, file.id)
		for i := 0; i < file.block_size; i++ {
			answer = append(answer, file.id)
		}
		for i := 0; i < file.block_free_space; i++ {
			answer = append(answer, -1)
		}
	}

	leftPos := 0
	rightPos := len(answer)-1
	
	// fmt.Println(answer)
	for leftPos < rightPos {
		if answer[leftPos] == -1 && answer[rightPos] != -1  {
			// fmt.Println("Start ", answer)
			answer[leftPos], answer[rightPos] = answer[rightPos], answer[leftPos]

		}
		for answer[leftPos] != -1 {
			leftPos ++
		} 
		for answer[rightPos] == -1 {
			rightPos --
		}
	}

	dot_found := false
	checkSum := 0
	for i, c := range(answer) {
		if c != -1 {
			checkSum += c*i
		} else if !dot_found {
			dot_found = true
			// fmt.Println()
			// fmt.Println(answer[i-3:i], answer[i:i+4])
			break
		}
	}

	return checkSum
}

func reGroup(answer [][]int) [][]int {
	var flatAnswer []int
	for _, a := range(answer) {
		flatAnswer = append(flatAnswer, a...)
	}

	// fmt.Println(flatAnswer)

	var newAnswer [][]int
	for i := 0; i < len(flatAnswer)-1; {
		var temp []int
		count := 0
		for i+count < len(flatAnswer) && flatAnswer[i] == flatAnswer[i + count] {
			if i+count < len(flatAnswer) {
				temp = append(temp, flatAnswer[i+count])
			}
			count++
		}
		newAnswer = append(newAnswer, temp)
		i+= count
	}
	// fmt.Println(newAnswer)
	return newAnswer
}