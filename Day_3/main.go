package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	
	P1()
	P2()
}

func P2() {
	file, err := os.Open("Day_3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	total := 0
	input:= ""
	do := true
	for scanner.Scan() {
		line := scanner.Text()
		// build input 
		for i, c := range(line) {
			if string(c) == "d" {
				if len(line) > i+len("do()") && line[i:i+len("do()")] == "do()" {
					do = true
				} else if len(line) > i+len("don't()") && line[i:i+len("don't()")] == "don't()" {
					do = false
				}
			} 
			if do {
				input += string(c)
			}
		}
	}	
	total += calcMatches(input)
	fmt.Println(total)
}
		
func P1() {
	file, err := os.Open("Day_3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += calcMatches(line)
	}
	fmt.Println(total)
}	

func findMatches(s string) []string {
	// Define the regex pattern
	pattern := `mul\(([0-9]{1,3}),([0-9]{1,3})\)` // Matches mul(a,b) where a and b are 1-3 digit numbers

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllString(s, -1) // -1 means no limit on the number of matches

	// Print matches
	// fmt.Println("Matches:", matches)
	return matches
}

func calcMatches(line string) int {
	total := 0
	matches := findMatches(line)
	for _, match := range(matches) {
		numbers_as_strings := strings.Split(match[4:len(match)-1], ",")
		d1, err := strconv.Atoi(numbers_as_strings[0])
		d2, err := strconv.Atoi(numbers_as_strings[1])
		
		if err != nil {
			// fmt.Println(err)
			} else {
			// fmt.Println(numbers_as_strings, d1, d2)
			total += d1*d2
		}
	}
	return total
}