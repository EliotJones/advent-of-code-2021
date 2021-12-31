package main

import (
	"bufio"
	"fmt"
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
	scanner, err := scannerForFile("day9.txt")

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
