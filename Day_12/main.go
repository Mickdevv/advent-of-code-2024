package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input [][]string
	// file, err:= os.Open("Day_12/input_test.txt")
	file, err:= os.Open("Day_12/input.txt")
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
	findRegions(input)
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
	total := 0
	for lineIndex, line := range(*input) {
		for charIndex := range line {{
			currentPos := [2]int{charIndex,lineIndex}
			_, exists := input_list_checked[currentPos]
			if !exists {
				fmt.Println("Checking region ")
				region := [][2]int{}
				fence := [][2]int{}
				region, fence = expandRegion1(input, &input_list_checked, currentPos, region, fence)
				for _, v := range(region) {
					fmt.Println("Recording checked values")
					input_list_checked[v] = getValueAtPosition(input, v)
				}
				region = removeDuplicates(region)
				regions = append(regions, region)
				fmt.Println("region", region)
				fmt.Println("fence ", fence)
				// fmt.Println("fence2", normalizeFence(input, fence, getValueAtPosition(input, region[0]), region))
				// fence = normalizeFence(input, fence, getValueAtPosition(input, region[0]), region)
				fences = append(fences, fence)
				fmt.Println(len(region), "*", len(fence), "=", len(region) * len(fence), getValueAtPosition(input, region[0]))
				total += len(region) * len(fence)
			}
		}
	}
	// for _, r := range(regions) {
		fmt.Println(total)
	}
}

func expandRegion1(input *[][]string, input_list_checked *map[[2]int]string, startingPos [2]int, region [][2]int, fence [][2]int) ([][2]int, [][2]int) {
	garden_plot := getValueAtPosition(input, startingPos)
	(*input_list_checked)[startingPos] = garden_plot
	potentialMoves, fence := getMovesAroundPositionWithFence(input, len((*input)[0]), len(*input), startingPos, input_list_checked)
	region = append(region, startingPos)
	for _, p := range(potentialMoves) {
		if getValueAtPosition(input, p) == garden_plot {
			region_append, fence_append := expandRegion1(input, input_list_checked, p, region, fence) 
			region, fence = append(region, region_append...), append(fence, fence_append...)
		}
	}
	return region, fence
}

func normalizeFence(input *[][]string, fence [][2]int, regionValue string, region [][2]int) [][2]int {
	fmt.Println("Normalizing fence")
	var normalizedFence [][2]int
	fence = removeDuplicates(fence)

	for _, pos := range(fence) {
		right := [2]int{pos[0] + 1, pos[1]}
		down := [2]int{pos[0], pos[1] + 1}
		left := [2]int{pos[0] - 1, pos[1]}
		up := [2]int{pos[0], pos[1] - 1}
		
		if getValueAtPosition(input, right) == regionValue && posInRegion(region, right) {
			normalizedFence = append(normalizedFence, pos)
		}
		if getValueAtPosition(input, left) == regionValue && posInRegion(region, left) {
			normalizedFence = append(normalizedFence, pos)
		}
		if getValueAtPosition(input, up) == regionValue && posInRegion(region, up) {
			normalizedFence = append(normalizedFence, pos)
		}
		if getValueAtPosition(input, down) == regionValue && posInRegion(region, down) {
			normalizedFence = append(normalizedFence, pos)
		}
	}
	return normalizedFence
}

func posInRegion(region [][2]int, pos [2]int) bool {
	for _, v := range(region) {
		if v == pos {
			return true
		}
	}
	return false
}

func findRegions(input *[][]string) {
	var input_list [][2]int
	// var regions [][][2]int
	fmt.Println("Importing data")
	for y:= 0; y < len((*input)); y++ {
		for x:= 0; x < len((*input)[0]); x++ {
			input_list = append(input_list, [2]int{x, y})
		}
	}
	regions := [][][2]int{}
	input_list_checked:= make(map[[2]int]string)
	
	fmt.Println("Processing data")
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
	fmt.Println("Expanding region")
	garden_plot := getValueAtPosition(input, startingPos)
	(*input_list_checked)[startingPos] = garden_plot
	fmt.Println("Getting moves around position")
	potentialMoves := getMovesAroundPosition(input, len((*input)[0]), len(*input), startingPos, input_list_checked)
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
	if pos[0] >= 0 && pos[1] >= 0 && pos[0] < len((*input)[0]) && pos[1] < len(*input) {
		return (*input)[pos[1]][pos[0]]
	}
	return ""
}

func getMovesAroundPositionWithFence(input *[][]string, w int, h int, pos [2]int, input_list_checked *map[[2]int]string) ([][2]int, [][2]int) {
	var positions [][2]int
	var fence [][2]int
	right := [2]int{pos[0] + 1, pos[1]}
	down := [2]int{pos[0], pos[1] + 1}
	left := [2]int{pos[0] - 1, pos[1]}
	up := [2]int{pos[0], pos[1] - 1}
	currentValue := getValueAtPosition(input, pos)

	if right[0] < w {
		_, exists := (*input_list_checked)[right]
		if !exists &&  getValueAtPosition(input, right) == currentValue {
			positions = append(positions, right)
		} else if getValueAtPosition(input, right) != currentValue {
			fence = append(fence, right)
		}
	} else {
		fence = append(fence, right)
	}
	if down[1] < h {
		_, exists := (*input_list_checked)[down]
		if !exists && getValueAtPosition(input, down) == currentValue {
			positions = append(positions, down)
		} else if getValueAtPosition(input, down) != currentValue {
			fence = append(fence, down)
		}	
	} else {
		fence = append(fence, down)
	}
	if left[0] >= 0 {
		_, exists := (*input_list_checked)[left]
		if !exists {
			positions = append(positions, left)
		} else if getValueAtPosition(input, left) != currentValue {
			fence = append(fence, left)
		}	
	} else {
		fence = append(fence, left)
	}
	if up[1] >= 0 {
		_, exists := (*input_list_checked)[up]
		if !exists && getValueAtPosition(input, up) != currentValue{
			positions = append(positions, up)
		} else {
			fence = append(fence, up)
		}	
	} else {
		fence = append(fence, up)
	}
	return positions, fence
}

func getMovesAroundPosition(input *[][]string, w int, h int, pos [2]int, input_list_checked *map[[2]int]string) [][2]int {
	var positions [][2]int
	right := [2]int{pos[0] + 1, pos[1]}
	down := [2]int{pos[0], pos[1] + 1}
	left := [2]int{pos[0] - 1, pos[1]}
	up := [2]int{pos[0], pos[1] - 1}
	currentValue := getValueAtPosition(input, pos)

	if right[0] < w {
		_, exists := (*input_list_checked)[right]
		if !exists &&  getValueAtPosition(input, right) == currentValue {
			positions = append(positions, right)
		} 
	} 
	if down[1] < h {
		_, exists := (*input_list_checked)[down]
		if !exists && getValueAtPosition(input, down) == currentValue {
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
		if !exists && getValueAtPosition(input, up) != currentValue{
			positions = append(positions, up)
		} 
	} 

	return positions
}

func removeDuplicates(list [][2]int) [][2]int {
	fmt.Println("Removing duplicates")
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


// func expandRegionIterative(input *[][]string, startingPos [2]int) [][2]int {
// 	var region [][2]int
// 	startingChar := getValueAtPosition(input, startingPos)
	
// 	return region
// }