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

func StringToInt(s string) int {
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