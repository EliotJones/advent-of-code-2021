package main

import (
	"bufio"
	"fmt"
	"sort"
)

func parseDay9Input(scanner *bufio.Scanner) [][]byte {
	const charToByteAdjustment = 48
	var grid [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]byte, len(line))
		for i := 0; i < len(line); i++ {
			b := line[i]
			row[i] = b - byte(charToByteAdjustment)
		}

		grid = append(grid, row)
	}

	return grid
}

func isLowestPoint(row int, col int, grid [][]byte) bool {
	val := grid[row][col]
	if val == 9 {
		return false
	}

	// Left.
	if col > 0 && grid[row][col-1] <= val {
		return false
	}

	// Right.
	if col < len(grid[row])-1 && grid[row][col+1] <= val {
		return false
	}

	// Top.
	if row > 0 && grid[row-1][col] <= val {
		return false
	}

	// Bottom.
	if row < len(grid)-1 && grid[row+1][col] <= val {
		return false
	}

	return true
}

func day9() {
	scanner, err := scannerForFile("inputs/day9.txt")

	if err != nil {
		panic(err)
	}

	grid := parseDay9Input(scanner)

	var totalRisk int
	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		for colIndex := 0; colIndex < len(grid[rowIndex]); colIndex++ {
			if isLowestPoint(rowIndex, colIndex, grid) {
				totalRisk += int(grid[rowIndex][colIndex]) + 1
			}
		}
	}

	fmt.Println("Result is", totalRisk)
}

func getBasinArea(location point, visited *[]point, grid [][]byte) int {
	row, col := location.y, location.x
	val := grid[row][col]

	// Stop condition.
	if val == 9 {
		return 0
	}

	// Stop if we've been here before.
	for _, p := range *visited {
		if p.x == location.x && p.y == location.y {
			return 0
		}
	}

	// This.
	result := 1
	*visited = append(*visited, location)
	if col > 0 {
		result += getBasinArea(*newPoint(col-1, row), visited, grid)
	}

	// Right.
	if col < len(grid[row])-1 {
		result += getBasinArea(*newPoint(col+1, row), visited, grid)
	}

	// Top.
	if row > 0 {
		result += getBasinArea(*newPoint(col, row-1), visited, grid)
	}

	// Bottom.
	if row < len(grid)-1 {
		result += getBasinArea(*newPoint(col, row+1), visited, grid)
	}

	return result
}

func day9p2() {
	scanner, err := scannerForFile("inputs/day9.txt")

	if err != nil {
		panic(err)
	}

	grid := parseDay9Input(scanner)

	var areas []int

	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		for colIndex := 0; colIndex < len(grid[rowIndex]); colIndex++ {
			if isLowestPoint(rowIndex, colIndex, grid) {
				var initial *[]point
				newSlice := make([]point, 0)
				initial = &newSlice

				area := getBasinArea(*newPoint(colIndex, rowIndex), initial, grid)
				areas = append(areas, area)
			}
		}
	}

	sort.Ints(areas)

	result := 1
	for i := len(areas) - 1; i >= len(areas)-3; i-- {
		result *= areas[i]
	}

	fmt.Println("Result is", result)
}
