package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type buttonConfig struct {
	A [2]float64;
	B [2]float64;
	Prize [2]float64;
}

func main() {
	// file, err := os.Open("Day_13/input_test.txt")
	file, err := os.Open("Day_13/input.txt")
	if err != nil {
		panic(err)
	}

	var input []string
	
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	P1(&input)
	P2(&input)
}

func P2(input *[]string) {
	buttonConfigs := extractButtonConfigs(input, 10000000000000)
	// fmt.Println(buttonConfigs)
	var possibleWins []buttonConfig
	total := 0
	// Filter out configurations with no solution
	for _, v := range(buttonConfigs) {
		a, b := calculateValue(v)
		if (a >= 0 && b >= 0) && (a > 0 || b > 0)  {
			possibleWins = append(possibleWins, v)
			total += calculateMoveCost(a, b)
		}
	}
	fmt.Println(total)
}

func P1(input *[]string) {
	buttonConfigs := extractButtonConfigs(input, 0)
	// fmt.Println(buttonConfigs)
	var possibleWins []buttonConfig
	total := 0
	// Filter out configurations with no solution
	for _, v := range(buttonConfigs) {
		a, b := calculateValue(v)
		if (a >= 0 && b >= 0) && (a > 0 || b > 0)  {
			possibleWins = append(possibleWins, v)
			total += calculateMoveCost(a, b)
		}
	}
	fmt.Println(total)
}


func calculateMoveCost(a int, b int) int {
	return int(a*3) + int(b)
}

func extractButtonConfigs(input *[]string, offset float64) []buttonConfig {
	var buttonConfigs []buttonConfig
	for i:= 0; i < len(*input); i+=4 {
		buttonAX := strToF32(strings.Split((*input)[i], " ")[2][2:len(strings.Split((*input)[i], " ")[2])-1])
		buttonAY := strToF32(strings.Split((*input)[i], " ")[3][2:])
		buttonBX := strToF32(strings.Split((*input)[i+1], " ")[2][2:len(strings.Split((*input)[i+1], " ")[2])-1])
		buttonBY := strToF32(strings.Split((*input)[i+1], " ")[3][2:])
		targetX := strToF32(strings.Split((*input)[i+2], " ")[1][2:len(strings.Split((*input)[i+2], " ")[1])-1]) + offset
		targetY := strToF32(strings.Split((*input)[i+2], " ")[2][2:]) + offset


		b := buttonConfig{A: [2]float64{buttonAX, buttonAY}, B: [2]float64{buttonBX, buttonBY}, Prize: [2]float64{targetX, targetY}, }
		// fmt.Println(b)
		buttonConfigs = append(buttonConfigs, b)
	}
	return buttonConfigs
}

func strToF32(input string) float64 {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return (float64(i))
}

func calculateValue(BC buttonConfig) (int, int) {
	a := (BC.Prize[0]*BC.B[1] - BC.Prize[1]*BC.B[0]) / (BC.A[0]*BC.B[1] - BC.A[1]*BC.B[0])
	b := (BC.A[0]*BC.Prize[1] - BC.A[1]*BC.Prize[0]) / (BC.A[0]*BC.B[1] - BC.A[1]*BC.B[0])

	if isInteger(a) && isInteger(b)  && a >= 0 && b >= 0 {
		return int(a), int(b)
	}
	return 0, 0
}

func abs(i float64) float64 {
	if i < 0 {
		return -i
	}
	return i
}

func isInteger(x float64) bool {
	return x == float64(int(x))
}