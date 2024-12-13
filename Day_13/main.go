package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type buttonConfig struct {
	A [2]string;
	B [2]string;
	Prize [2]string;
}

func main() {
	file, err := os.Open("Day_13/input_test.txt")
	if err != nil {
		panic(err)
	}

	var input []string
	
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	// P2(&input)
	P1(&input)
}

func P2(input *[]string) {
}

func P1(input *[]string) {
	fmt.Println(extractButtonConfigs(input))
	
}

func extractButtonConfigs(input *[]string) []buttonConfig {
	var buttonConfigs []buttonConfig
	for i:= 0; i < len(*input); i+=4 {
		buttonAX := strings.Split((*input)[i], " ")[2][2:len(strings.Split((*input)[i], " ")[2])-1]
		buttonAY := strings.Split((*input)[i], " ")[2][2:]
		buttonBX := strings.Split((*input)[i+1], " ")[2][2:len(strings.Split((*input)[i+1], " ")[2])-1]
		buttonBY := strings.Split((*input)[i+1], " ")[2][2:]
		targetX := strings.Split((*input)[i+2], " ")[1][2:len(strings.Split((*input)[i+2], " ")[2])-1]
		targetY := strings.Split((*input)[i+2], " ")[1]
		b := buttonConfig{A: [2]string{buttonAX, buttonAY}, B: [2]string{buttonBX, buttonBY}, Prize: [2]string{targetX, targetY}, }
		fmt.Println(b)
		buttonConfigs = append(buttonConfigs, b)
	}
	return buttonConfigs
}