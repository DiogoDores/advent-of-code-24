package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFile(fileName string) []string {
	file, err := os.Open(fileName)
	Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	Check(scanner.Err())
	return lines
}

// Check panics if error is nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Stoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func CreateMatrix(data []string) [][]string {
	var matrix [][]string
	for _, line := range data {
		var row []string
		for _, c := range line {
			row = append(row, string(c))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func TransposeMatrix(matrix [][]string) [][]string {
	var transposedMatrix [][]string
	
	for i := 0; i < len(matrix[0]); i++ {
		var row []string
		for j := 0; j < len(matrix); j++ {
			row = append(row, matrix[j][i])
		}
		transposedMatrix = append(transposedMatrix, row)
	}

	return transposedMatrix
}


func ReverseString(str string) (result string) { 
    for _, v := range str { 
        result = string(v) + result 
    } 
    return
} 