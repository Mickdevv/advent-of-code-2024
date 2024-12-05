package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var table []string
	file, err := os.Open("Day_5/input_test.txt")
	if err!= nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		table = append(table, line)
	}
	P1(table)
	P2(table)
}

func P2(table []string) {
	fmt.Println(table)
}

func P1(table []string) {
	fmt.Println(table)
}