package main

import (
	"fmt"
	"slices"
	"strings"
	"utils"
)

func readAndSortArrays(data []string) ([]int, []int) {
	leftside, rightside := []int{}, []int{}
	for _, line := range data {
		splitData := strings.Split(line, "   ")
		leftValue, rightValue := utils.Stoi(splitData[0]), utils.Stoi(splitData[1])
		leftside = append(leftside, leftValue)
		rightside = append(rightside, rightValue)
	}

	slices.Sort(leftside)
	slices.Sort(rightside)

	return leftside, rightside
}

func part1(leftside, rightside []int) int {
	sum := 0
	for i, left := range leftside {
		sum += utils.Abs(left - rightside[i])
	}

	return sum
}

func part2(leftside, rightside []int) int {
	sum := 0
	counts := make(map[int]int)
	for _, left := range leftside {
		counts[left] = 0
		for _, right := range rightside {
			if right == left {
				counts[left]++
			}
		}
		sum += left * counts[left]
	}
	return sum
}

func main() {
	data := utils.ReadFile("input.txt")
	leftside, rightside := readAndSortArrays(data)

	sumPart1 := part1(leftside, rightside)
	sumPart2 := part2(leftside, rightside)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", sumPart1, sumPart2)
}
