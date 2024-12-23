package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	connections := make(map[string][]string)
	sets := [][3]string{}
	file, err := os.Open("Day_23/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")

		left, right := line[0], line[1]

		_, leftExists := connections[left]
		if leftExists {
			connections[left] = append(connections[left], right)
		} else {
			connections[left] = []string{right}
		}

		_, rightExists := connections[right]
		if rightExists {
			connections[right] = append(connections[right], left)
		} else {
			connections[right] = []string{left}
		}
	}

	fmt.Println(connections)
	for k, v := range(connections) {
		for _, c2 := range(v) {
			for _, c3 := range(connections[c2]) {
				if isStringInSlice(c3, v) {
					sets = append(sets, [3]string{k, c2, c3})
				}
			}
		}
	}
	stringSets := filterSets(sets)
	fmt.Println(len(stringSets))
}

func isStringInSlice(s string, sl []string) bool {
	for _, v := range(sl) {
		if string(v) == s {
			return true
		}
	}
	return false
}

func filterSets(sets [][3]string) []string {
	count := 0
	newSetmap := make(map[string]bool)
	for _, set := range(sets) {
		if checkForT(set) {
			count ++
			// fmt.Println("=====")
			// fmt.Println(set)
			// fmt.Println(normalizeSlice(set[:]))
			newSetmap[normalizeSlice(set[:])] = true
			// fmt.Println(newSetmap)
		}
	}
	var newSetArray []string
	for k := range newSetmap {
		newSetArray = append(newSetArray, k)
		fmt.Println(k)
	}
	return newSetArray
}

func checkForT(s [3]string) bool {
	for _, v := range(s) {
		if strings.Contains(strings.Split(string(v), "")[0], "t") {
			return true
		}
	}
	return false
}

func normalizeSlice(slice []string) string {
	sort.Strings(slice) // Sort the slice to ensure consistent order
	return strings.Join(slice, ",")
}