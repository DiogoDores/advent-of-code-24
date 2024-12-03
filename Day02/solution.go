package main

import (
	"fmt"
	"strings"
	"utils"
)

func isSafe(report []string, ascending bool, problemDampener bool) bool {
	if checkSafety(report, ascending) {
		return true
	}
	if problemDampener {
		for i := 0; i < len(report); i++ {
			modifiedReport := utils.RemoveIndex(report, i)
			if checkSafety(modifiedReport, ascending) {
				return true
			}
		}
	}
	return false
}

func checkSafety(report []string, ascending bool) bool {
	oldLevel := utils.Stoi(report[0])
	for i := 1; i < len(report); i++ {
		newLevel := utils.Stoi(report[i])
		delta := utils.Abs(newLevel - oldLevel)
		if (ascending && newLevel < oldLevel) || (!ascending && newLevel > oldLevel) || delta > 3 || delta < 1 {
			return false
		}
		oldLevel = newLevel
	}
	return true
}

func part1(report []string) int {
	if isSafe(report, true, false) || isSafe(report, false, false) {
		return 1
	}
	return 0
}

func part2(report []string) int {
	if isSafe(report, true, true) || isSafe(report, false, true) {
		return 1
	}
	return 0
}

func main() {
	data := utils.ReadFile("input.txt")
	safeCountPart1, safeCountPart2 := 0, 0

	for _, line := range data {
		report := strings.Split(line, " ")
		safeCountPart1 += part1(report)
		safeCountPart2 += part2(report)
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", safeCountPart1, safeCountPart2)
}
