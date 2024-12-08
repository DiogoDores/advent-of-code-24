package main

import (
	"fmt"
	"strings"
	"utils"
	"math"
)

func evaluateExpression(nums []int, operators []string) int {
	result := nums[0]
	for i, op := range operators {
		switch op {
			case "+":
				result += nums[i+1]
			case "*":
				result *= nums[i+1]
			case "||":
				result = utils.Stoi(utils.Itos(result) + utils.Itos(nums[i+1]))
		}
	}
	return result
}

func isValidEquation(result int, nums []int, concat bool) bool {
	n := len(nums) - 1
	rng := int(math.Pow(2, float64(n)))
	if concat {
		rng = int(math.Pow(3, float64(n)))
	}

	for i := 0; i < rng; i++ {
		operators := make([]string, n)
		for j := 0; j < n; j++ {
			generateOperators(i, j, operators, concat)
		}
		if evaluateExpression(nums, operators) == result {
			return true
		}
	}
	return false
}

func generateOperators(i int, j int, operators []string, concat bool) {
	if concat {
		switch (i / int(math.Pow(3, float64(j)))) % 3 {
			case 0:
				operators[j] = "+"
			case 1:
				operators[j] = "*"
			case 2:
				operators[j] = "||"
		}
	} else {
		if (i / int(math.Pow(2, float64(j)))) % 2 != 0 {
			operators[j] = "*"
		} else {
			operators[j] = "+"
		}
	}
}

func main() {
	data := utils.ReadFile("input.txt")
	totalSum1, totalSum2 := 0, 0

	for _, line := range data {
		parts := strings.Split(line, ":")
		result := utils.Stoi(parts[0])
		nums := utils.SliceStoi(strings.Fields(parts[1]))

		if isValidEquation(result, nums, false) {
			totalSum1 += result
		}

		if isValidEquation(result, nums, true) {
			totalSum2 += result
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", totalSum1, totalSum2)
}