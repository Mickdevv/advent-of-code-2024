package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("Day_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	P2(scanner)

}

func P2(scanner *bufio.Scanner) {
	left_numbers := []int{}
	right_numbers := []int{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		
		left_numbers = append(left_numbers, string_to_int(line[0]))
		right_numbers = append(right_numbers, string_to_int(line[len(line)-1]))
	}
	left_numbers = bubbleSort(left_numbers)
	right_numbers = bubbleSort(right_numbers)

	total := 0
	checked_numbers := []int{}
	for i := range(len(left_numbers)) {
		if contains(checked_numbers, left_numbers[i]) == false {
			checked_numbers = append(checked_numbers, left_numbers[i])
			total += left_numbers[i]*countOccurrences(left_numbers[i], right_numbers)
		}
	}

	fmt.Println(total)
}

func P1(scanner *bufio.Scanner) {
	left_numbers := []int{}
	right_numbers := []int{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		
		left_numbers = append(left_numbers, string_to_int(line[0]))
		right_numbers = append(right_numbers, string_to_int(line[len(line)-1]))
	}
	// fmt.Println(left_numbers)
	left_numbers = bubbleSort(left_numbers)
	// fmt.Println(left_numbers)
	// fmt.Println(right_numbers)
	right_numbers = bubbleSort(right_numbers)
	// fmt.Println(right_numbers)

	total := 0

	check_left := 0
	check_right := 0
	for i := range(len(left_numbers)) {
		if check_left > left_numbers[i] {
			fmt.Println("Check left", i, check_left, left_numbers[i])
		}
		if check_right > right_numbers[i] {
			fmt.Println("Check right", i, check_right, right_numbers[i])
		}
		// fmt.Println(left_numbers[i], right_numbers[i], abs(left_numbers[i] - right_numbers[i]))
		total += abs(left_numbers[i] - right_numbers[i])
		check_left = left_numbers[i]
		check_right = right_numbers[i]
	}

	fmt.Println(total)
}

func string_to_int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func bubbleSort(arr []int) []int {
	for i := range(len(arr)) {
		for j := range(len(arr)) {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func countOccurrences(i int, arr []int) int {
	count := 0
	for _, v := range(arr) {
		if v == i {
			count ++
		}
	}
	return count
}

func contains(arr []int, i int) bool {
	for _, v := range(arr) {
		if v == i {
			return true
		}
	}
	return false
}