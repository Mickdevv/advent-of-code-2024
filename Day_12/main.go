package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input [][]string
	file, err:= os.Open("Day_12/input_test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, strings.Split(line, ""))
	}

	P1(&input)
}

func P1(input *[][]string) {
	// var regions [][][2]int
	findRegions1(input)
}

func P2(input *[][]string) {
	// fmt.Println(*input)
}

func findRegions1(input *[][]string) {
	var input_list [][2]int
	// var regions [][][2]int
	for y:= 0; y < len((*input)); y++ {
		for x:= 0; x < len((*input)[0]); x++ {
			input_list = append(input_list, [2]int{x, y})
		}
	}
	regions := [][][2]int{}
	fences := [][][2]int{}
	input_list_checked:= make(map[[2]int]string)

	for lineIndex, line := range(*input) {
		for charIndex := range line {{
			currentPos := [2]int{charIndex,lineIndex}
			_, exists := input_list_checked[currentPos]
			if !exists {
				region := [][2]int{}
				fence := [][2]int{}
				region, fence = expandRegion1(input, &input_list_checked, currentPos, region, fence)
				for _, v := range(region) {
					input_list_checked[v] = getValueAtPosition(input, v)
				}
				regions = append(regions, removeDuplicates(region))
				fences = append(fences, removeDuplicates(fence))
				fmt.Println(len(region), len(fence))
			}
		}
	}
	// for _, r := range(regions) {
	// 	fmt.Println(r)
	}
}

func expandRegion1(input *[][]string, input_list_checked *map[[2]int]string, startingPos [2]int, region [][2]int, fence [][2]int) ([][2]int, [][2]int) {
	garden_plot := getValueAtPosition(input, startingPos)
	(*input_list_checked)[startingPos] = garden_plot
	potentialMoves := getMovesAroundPosition(len((*input)[0]), len(*input), startingPos, input_list_checked)
	region = append(region, startingPos)
	for _, p := range(potentialMoves) {
		if getValueAtPosition(input, p) == garden_plot {
			region_append, fence_append := expandRegion1(input, input_list_checked, p, region, fence) 
			region, fence = append(region, region_append...), append(fence, fence_append...)
		}
	}
	return region, fence
}

func findRegions(input *[][]string) {
	var input_list [][2]int
	// var regions [][][2]int
	for y:= 0; y < len((*input)); y++ {
		for x:= 0; x < len((*input)[0]); x++ {
			input_list = append(input_list, [2]int{x, y})
		}
	}
	regions := [][][2]int{}
	input_list_checked:= make(map[[2]int]string)
	
	for lineIndex, line := range(*input) {
		for charIndex := range(line) {
			region := [][2]int{}
			currentPos := [2]int{charIndex,lineIndex}
			_, exists := input_list_checked[currentPos]
			if !exists {
				region = removeDuplicates(expandRegion(input, &input_list_checked, currentPos, region))
				for _, v := range(region) {
					input_list_checked[v] = getValueAtPosition(input, v)
				}
				regions = append(regions, region)
			}
		}
	}
	for _, r := range(regions) {
		fmt.Println(r)
	}
}

func expandRegion(input *[][]string, input_list_checked *map[[2]int]string, startingPos [2]int, region [][2]int) [][2]int {
	garden_plot := getValueAtPosition(input, startingPos)
	(*input_list_checked)[startingPos] = garden_plot
	potentialMoves := getMovesAroundPosition(len((*input)[0]), len(*input), startingPos, input_list_checked)
	region = append(region, startingPos)
	if len(potentialMoves) == 0 {
		return region
	}
	for _, p := range(potentialMoves) {
		if getValueAtPosition(input, p) == garden_plot {
			region = append(region, expandRegion(input, input_list_checked, p, region)...) 
		}
	}
	return region
}

func getValueAtPosition(input *[][]string, pos [2]int) string {
	return (*input)[pos[1]][pos[0]]
}

func getMovesAroundPosition(w int, h int, pos [2]int, input_list_checked *map[[2]int]string) [][2]int {
	var positions [][2]int
	right := [2]int{pos[0] + 1, pos[1]}
	down := [2]int{pos[0], pos[1] + 1}
	left := [2]int{pos[0] - 1, pos[1]}
	up := [2]int{pos[0], pos[1] - 1}

	if right[0] < w {
		_, exists := (*input_list_checked)[right]
		if !exists {
			positions = append(positions, right)
		}
	}
	if down[1] < h {
		_, exists := (*input_list_checked)[down]
		if !exists {
			positions = append(positions, down)
		}	
	}
	if left[0] >= 0 {
		_, exists := (*input_list_checked)[left]
		if !exists {
			positions = append(positions, left)
		}	
	}
	if up[1] >= 0 {
		_, exists := (*input_list_checked)[up]
		if !exists {
			positions = append(positions, up)
		}	
	}
	return positions
}

func removeDuplicates(list [][2]int) [][2]int {
	listMap := make(map[[2]int][2]int)
	var newList [][2]int
	for _, v := range(list) {
		listMap[v] = v
	}
	for _, v := range(listMap) {
		newList = append(newList, v)
	}
	return newList
}
