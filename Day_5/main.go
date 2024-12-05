package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var table []string
	var rules [][]string
	var updates [][]string
	file, err := os.Open("Day_5/input.txt")
	if err!= nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		
		line := scanner.Text()
		table = append(table, line)
		if strContains(line, ",") {
			updates = append(updates, strings.Split(line, ","))
		} else if strContains(line, "|") {
			rules = append(rules, strings.Split(line, "|"))
		}
	}
	P1(rules, updates)
	P2(rules, updates)
}

func P1(rules [][]string, updates [][]string) {
// 	fmt.Println(rules)
// 	fmt.Println(updates)
	
	total:=0
	for _, update := range(updates) {
		ordered := true
		for _, rule := range(rules) {
			if arrContains(update, rule[0]) != -1 && arrContains(update, rule[1]) != -1 {
				if arrContains(update, rule[0]) > arrContains(update, rule[1]) {
					// fmt.Println(update, rule, arrContains(update, rule[0]), arrContains(update, rule[1]))
					ordered = false
					break
				}
			}
		}
		if ordered {
			middleValue, _ := strconv.Atoi(update[(len(update) -1)/2])
			total += middleValue
			// fmt.Println(update, middleValue)
		}
	}
	fmt.Println(total)
}

func P2(rules [][]string, updates [][]string) {
// 	fmt.Println(rules)
// 	fmt.Println(updates)
	
	total:=0
	for _, update := range(updates) {
		ordered := true
		for !isCorrectlyOrdered(update, rules) {
			for _, rule := range(rules) {
				if arrContains(update, rule[0]) != -1 && arrContains(update, rule[1]) != -1 {
					if arrContains(update, rule[0]) > arrContains(update, rule[1]) {
						// fmt.Println()
						// fmt.Println(update, rule, arrContains(update, rule[0]), arrContains(update, rule[1]))
						tempValue := update[arrContains(update, rule[0])]
						update[arrContains(update, rule[0])] = update[arrContains(update, rule[1])]
						update[arrContains(update, rule[1])] = tempValue
						ordered = false
						// fmt.Println(update, rule, arrContains(update, rule[0]), arrContains(update, rule[1]))
					}
				}
			}
		}
		if !ordered {
			middleValue, _ := strconv.Atoi(update[(len(update) -1)/2])
			total += middleValue
			// fmt.Println(update, middleValue)
		}
	}
	fmt.Println(total)
}


func strContains(s string, char string) bool {
	for _, c := range s {
		if char == string(c) {
			return true
		}
	}
	return false
}

func arrContains(s []string, item string) int {
	for i, v := range s {
		if item == string(v) {
			return i
		}
	}
	return -1
}

func isCorrectlyOrdered(update []string, rules [][]string) bool {
	for _, rule := range(rules) {
		if arrContains(update, rule[0]) != -1 && arrContains(update, rule[1]) != -1 {
			if arrContains(update, rule[0]) > arrContains(update, rule[1]) {
				return false
			}
		}
	}
	return true
}