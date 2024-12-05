package main

import (
    "fmt"
    "strings"
    "utils"
)

func wordCount(line string) int {
    return strings.Count(line, "XMAS")
}

func checkDiagonal(matrix [][]string, i, j, di, dj int) int {
    word := "XMAS"
    for k := 1; k < len(word); k++ {
        i += di
        j += dj
        if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) || matrix[i][j] != string(word[k]) {
            return 0
        }
    }
    return 1
}

func countDiagonals(matrix [][]string) int {
    count := 0
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == "X" {
                count += checkDiagonal(matrix, i, j, 1, 1)
                count += checkDiagonal(matrix, i, j, 1, -1)
                count += checkDiagonal(matrix, i, j, -1, 1)
                count += checkDiagonal(matrix, i, j, -1, -1)
            }
        }
    }
    return count
}

func part1(matrix [][]string) int {
    transposedMatrix := utils.TransposeMatrix(matrix)
    count := 0

    for _, row := range matrix {
        line := strings.Join(row, "")
        count += wordCount(line)
        count += wordCount(utils.ReverseString(line))
    }

    for _, row := range transposedMatrix {
        line := strings.Join(row, "")
        count += wordCount(line)
        count += wordCount(utils.ReverseString(line))
    }

    count += countDiagonals(matrix)
    return count
}

func checkXMas(matrix [][]string, i, j int) int {
    coords := [][]int{
        {i - 1, j - 1}, {i - 1, j + 1},
        {i + 1, j - 1}, {i + 1, j + 1},
    }

    for _, coord := range coords {
        if coord[0] < 0 || coord[1] < 0 || coord[0] >= len(matrix) || coord[1] >= len(matrix[0]) {
            return 0
        }
    }

    if (matrix[coords[0][0]][coords[0][1]] == "M" && matrix[coords[3][0]][coords[3][1]] == "S") ||
        (matrix[coords[0][0]][coords[0][1]] == "S" && matrix[coords[3][0]][coords[3][1]] == "M") {
        if (matrix[coords[1][0]][coords[1][1]] == "M" && matrix[coords[2][0]][coords[2][1]] == "S") ||
            (matrix[coords[1][0]][coords[1][1]] == "S" && matrix[coords[2][0]][coords[2][1]] == "M") {
            return 1
        }
    }
    return 0
}

func part2(matrix [][]string) int {
    count := 0
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == "A" {
                count += checkXMas(matrix, i, j)
            }
        }
    }
    return count
}

func main() {
    data := utils.ReadFile("input.txt")
    matrix := utils.CreateMatrix(data)

    fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(matrix), part2(matrix))
}