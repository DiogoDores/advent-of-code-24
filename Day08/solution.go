package main

import (
	"fmt"
	"utils"
)

type coord struct {
	i int
	j int
}

func findFrequencies(matrix [][]string) map[string][]coord {
	frequencies := make(map[string][]coord)
	for i, row := range matrix {
		for j, cell := range row {
			if cell != "." {
				frequencies[cell] = append(frequencies[cell], coord{i, j})
			}
		}
	}
	return frequencies
}

func findAntinode(coord1 coord, coord2 coord, matrix [][]string) ([]coord, []coord) {
	var antinodes []coord
	var antinodesWithHarmonics []coord
	
	diffI, diffJ := utils.Abs(coord1.i - coord2.i), utils.Abs(coord1.j - coord2.j)
	lowestI, lowestJ := min(coord1.i, coord2.i), min(coord1.j, coord2.j)
	highestI, highestJ := max(coord1.i, coord2.i), max(coord1.j, coord2.j)
	antinode1, antinode2 := coord{}, coord{}
	antinodesWithHarmonics = append(antinodesWithHarmonics, coord1, coord2)

	if coord1.j > coord2.j {
		antinode1 = coord{lowestI - diffI, highestJ + diffJ}
		antinode2 = coord{highestI + diffI, lowestJ - diffJ}
		if utils.IsCoordInsideMatrix(antinode1.i, antinode1.j, matrix) {
			antinodes = append(antinodes, antinode1)
			antinodesWithHarmonics = append(antinodesWithHarmonics, antinode1, coord1, coord2)
			for {
				antinode1.i -= diffI
				antinode1.j += diffJ
				if !utils.IsCoordInsideMatrix(antinode1.i, antinode1.j, matrix) {
					break
				}
				antinodesWithHarmonics = append(antinodesWithHarmonics, antinode1)
			}
		}

		if utils.IsCoordInsideMatrix(antinode2.i, antinode2.j, matrix) {
			antinodes = append(antinodes, antinode2)
			antinodesWithHarmonics = append(antinodesWithHarmonics, antinode2)
			for {
				antinode2.i += diffI
				antinode2.j -= diffJ
				if !utils.IsCoordInsideMatrix(antinode2.i, antinode2.j, matrix) {
					break
				}
				antinodesWithHarmonics = append(antinodesWithHarmonics, antinode2)
			}
		}
	} else {
		antinode1 = coord{lowestI - diffI, lowestJ - diffJ}
		antinode2 = coord{highestI + diffI, highestJ + diffJ}

		if utils.IsCoordInsideMatrix(antinode1.i, antinode1.j, matrix) {
			antinodes = append(antinodes, antinode1)
			antinodesWithHarmonics = append(antinodesWithHarmonics, antinode1)

			for {
				antinode1.i -= diffI
				antinode1.j -= diffJ
				if !utils.IsCoordInsideMatrix(antinode1.i, antinode1.j, matrix) {
					break
				}
				antinodesWithHarmonics = append(antinodesWithHarmonics, antinode1)
			}
		}

		if utils.IsCoordInsideMatrix(antinode2.i, antinode2.j, matrix) {
			antinodes = append(antinodes, antinode2)
			antinodesWithHarmonics = append(antinodesWithHarmonics, antinode2)
			for {
				antinode2.i += diffI
				antinode2.j += diffJ
				if !utils.IsCoordInsideMatrix(antinode2.i, antinode2.j, matrix) {
					break
				}
				antinodesWithHarmonics = append(antinodesWithHarmonics, antinode2)
			}
		}
	}

	return antinodes, antinodesWithHarmonics
}

func findAllAntinodes(frequencies map[string][]coord, matrix [][]string) ([]coord, []coord) {
	var allAntinodes []coord
	var allAntinodesWithHarmonics []coord
	for _, freqList := range frequencies {
		counter := 0
		for counter < len(freqList) - 1 {
			for i := counter + 1; i < len(freqList); i++ {
				antinodes, antinodesWithHarmonics := findAntinode(freqList[counter], freqList[i], matrix)
				allAntinodes = append(allAntinodes, antinodes...)
				allAntinodesWithHarmonics = append(allAntinodesWithHarmonics, antinodesWithHarmonics...)
			}
			counter++
		}
	}
	return allAntinodes, allAntinodesWithHarmonics
}

func main() {
	data := utils.ReadFile("input.txt")
	matrix := utils.CreateMatrix(data)

	frequencies := findFrequencies(matrix)
	allAntinodes, allAntinodesWithHarmonics := findAllAntinodes(frequencies, matrix)
	allAntinodes = utils.RemoveDuplicate(allAntinodes)
	allAntinodesWithHarmonics = utils.RemoveDuplicate(allAntinodesWithHarmonics) 
	
	fmt.Println("Part 1:", len(allAntinodes), "\nPart 2:", len(allAntinodesWithHarmonics))
}