package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("Day_4/input_test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}