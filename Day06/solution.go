package main

import (
	"fmt"
	"utils"
)

type Guard struct {
	i, j      int
	direction string
}

type State struct {
	i, j int
	dir  string
}

var directions = map[string][2]int{
	"up":    {-1, 0},
	"right": {0, 1},
	"down":  {1, 0},
	"left":  {0, -1},
}

var dirOrder = []string{"up", "right", "down", "left"}

func findGuard(matrix [][]string) Guard {
	for i, row := range matrix {
		for j, cell := range row {
			if cell == "^" {
				return Guard{i, j, "up"}
			}
		}
	}
	return Guard{-1, -1, ""}
}

func getNextDirection(currentDirection string) string {
	for i, dir := range dirOrder {
		if dir == currentDirection {
			return dirOrder[(i+1)%4]
		}
	}
	return ""
}

func moveGuard(guard *Guard, matrix [][]string) {
	for {
		newI := guard.i + directions[guard.direction][0]
		newJ := guard.j + directions[guard.direction][1]

		if newI < 0 || newI >= len(matrix) || newJ < 0 || newJ >= len(matrix[0]) {
			guard.i, guard.j = -1, -1
			return
		}
		if matrix[newI][newJ] == "#" {
			guard.direction = getNextDirection(guard.direction)
		} else {
			guard.i, guard.j = newI, newJ
			return
		}
	}
}

func detectLoop(guard Guard, matrix [][]string) bool {
	visited := make(map[State]bool)

	for guard.i != -1 && guard.j != -1 {
		state := State{guard.i, guard.j, guard.direction}
		if visited[state] {
			return true
		}
		visited[state] = true
		moveGuard(&guard, matrix)
	}
	return false
}

func findObstructionPoints(guard Guard, matrix [][]string) []string {
	var obstructionPoints []string

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "." {
				matrix[i][j] = "#"
				if detectLoop(guard, matrix) {
					obstructionPoints = append(obstructionPoints, fmt.Sprintf("%d,%d", i, j))
				}
				matrix[i][j] = "."
			}
		}
	}
	return obstructionPoints
}

func main() {
	lines := utils.ReadFile("input.txt")
	matrix := utils.CreateMatrix(lines)

	guard := findGuard(matrix)
	initialPath := []string{}
	visited := make(map[string]bool)

	for guard.i != -1 && guard.j != -1 {
		position := fmt.Sprintf("%d,%d", guard.i, guard.j)
		if !visited[position] {
			initialPath = append(initialPath, position)
		}
		visited[position] = true
		moveGuard(&guard, matrix)
	}

	obstructionPoints := findObstructionPoints(findGuard(matrix), matrix)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", len(initialPath), len(obstructionPoints))
}
