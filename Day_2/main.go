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
	
	P1()
	P2()
}


func P1() {
	// file, err := os.Open("input.txt")
	file, err := os.Open("Day_2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	safe_reports := 0

	for scanner.Scan() {
		line:=strings.Split(scanner.Text(), " ")
		
		if isReportSafe(line) == true {
			safe_reports++
		}

	}
	fmt.Println(safe_reports)
}

func P2() {
		file, err := os.Open("Day_2/input.txt")
		// file, err := os.Open("input_test.txt")
		if err != nil {
			log.Fatal(err)
		}
	
		scanner := bufio.NewScanner(file)
		safe_reports := 0
	
		for scanner.Scan() {
			line:=strings.Split(scanner.Text(), " ")

			for j := 0; j < len(line); j++ {
				if isReportSafe(line) {
					safe_reports++
					break
				}
				// fmt.Println("Before append:", line, "j:", j)
				test_line := append([]string(nil), line[:j]...) // create a new slice from line[:j]
				test_line = append(test_line, line[j+1:]...) // append line[j+1:]        fmt.Println("test_line:", test_line, "original line:", line)
				if isReportSafe(test_line) {
					// fmt.Println("Line:      ", line)
					// fmt.Println("Test line: ", test_line, "j:", j)
					safe_reports++
					// fmt.Println()
					break
				}
			}
	
		}
		fmt.Println(safe_reports)
}

func string_to_int(str string) float64 {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return float64(i)
}

func isReportSafe(line []string) bool {
	safe := true
	if string_to_int(line[0]) > string_to_int(line[1]) {
		for i := 0; i < len(line)-1; i++ {
			if string_to_int(line[i]) - string_to_int(line[i+1]) > 3 || string_to_int(line[i]) - string_to_int(line[i+1]) < 1 {
				safe = false
				// fmt.Println(string_to_int(line[i]), string_to_int(line[i+1]), "increasing")
				break
			}
		}

	//Ascending
	} else {
		for i := 0; i < len(line)-1; i++ {
			if string_to_int(line[i+1]) - string_to_int(line[i]) > 3 || string_to_int(line[i+1]) - string_to_int(line[i]) < 1 {
				// fmt.Println(string_to_int(line[i]), string_to_int(line[i+1]), "decreasing")
				safe = false
				break
			}
		}
	}
	return safe
}