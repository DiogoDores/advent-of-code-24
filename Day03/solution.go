package main

import (
    "fmt"
    "regexp"
    "strings"
    "utils"
)

func calculateMul(mul string) int {
    regex := regexp.MustCompile(`\d{1,3}`)
    matches := regex.FindAllString(mul, -1)
    num1 := utils.Stoi(matches[0])
    num2 := utils.Stoi(matches[1])
    return num1 * num2
}

func executeInstructions(instructs []string) []string {
    do := true
    var finalInstructions []string

    for _, instruct := range instructs {
        switch instruct {
			case "do()":
				do = true
			case "don't()":
				do = false
			default:
				if do {
					finalInstructions = append(finalInstructions, instruct)
				}
        }
    }
    return finalInstructions
}

func main() {
    memory := strings.Join(utils.ReadFile("input.txt"), "")
    regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
    regexWithInstr := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)
    part1Mults, part2Mults := 0, 0

    for _, match := range regex.FindAllString(memory, -1) {
        part1Mults += calculateMul(match)
    }

    for _, match := range executeInstructions(regexWithInstr.FindAllString(memory, -1)) {
        part2Mults += calculateMul(match)
    }

    fmt.Printf("Part 1: %d\nPart 2: %d\n", part1Mults, part2Mults)
}