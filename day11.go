package main

import (
	"fmt"
	"math"
)

func parseFileToGrid(file string) [][]byte {
	scanner, err := scannerForFile(file)
	if err != nil {
		panic(err)
	}

	var result [][]byte

	for scanner.Scan() {
		line := scanner.Text()

		lineResult := make([]byte, len(line))

		for i := 0; i < len(line); i++ {
			b := line[i]
			lineResult[i] = b - byte(asciiNumToByteAdjustment)
		}

		result = append(result, lineResult)
	}

	return result
}

func incrementAtIndex(grid *[][]byte, rowIndex int, colIndex int) {
	val := (*grid)[rowIndex][colIndex]
	if val > 0 {
		(*grid)[rowIndex][colIndex]++
	}
}

func incrementNeighbours(grid *[][]byte, rowIndex int, colIndex int) {
	// Top row.
	if rowIndex > 0 {
		prevRow := (*grid)[rowIndex-1]

		// Above.
		incrementAtIndex(grid, rowIndex-1, colIndex)

		// Above left.
		if colIndex > 0 {
			incrementAtIndex(grid, rowIndex-1, colIndex-1)
		}

		// Above right.
		if colIndex < len(prevRow)-1 {
			incrementAtIndex(grid, rowIndex-1, colIndex+1)
		}
	}

	// This row.
	row := (*grid)[rowIndex]

	// Left.
	if colIndex > 0 {
		incrementAtIndex(grid, rowIndex, colIndex-1)
	}

	// Right.
	if colIndex < len(row)-1 {
		incrementAtIndex(grid, rowIndex, colIndex+1)
	}

	// Below.
	if rowIndex < len(*grid)-1 {
		below := (*grid)[rowIndex+1]

		incrementAtIndex(grid, rowIndex+1, colIndex)

		if colIndex > 0 {
			incrementAtIndex(grid, rowIndex+1, colIndex-1)
		}

		if colIndex < len(below)-1 {
			incrementAtIndex(grid, rowIndex+1, colIndex+1)
		}
	}
}

func incrementAllValuesByOne(grid *[][]byte) {
	for rowIndex := 0; rowIndex < len(*grid); rowIndex++ {
		row := (*grid)[rowIndex]
		for colIndex := 0; colIndex < len(row); colIndex++ {
			row[colIndex]++
		}
	}
}

func incrementDay11GridCountingFlashes(grid *[][]byte) int {
	var flashes int
	incrementAllValuesByOne(grid)

	// Could do something recursive but let's just brute force.
	for {
		var flashesThisRun int
		for rowIndex := 0; rowIndex < len(*grid); rowIndex++ {
			row := (*grid)[rowIndex]
			for colIndex := 0; colIndex < len(row); colIndex++ {
				val := row[colIndex]
				if val > 9 {
					row[colIndex] = 0

					flashes++
					flashesThisRun++

					incrementNeighbours(grid, rowIndex, colIndex)
				}
			}
		}

		if flashesThisRun == 0 {
			break
		}
	}

	return flashes
}

func incrementDay11GridFindSyncPoint(grid *[][]byte) bool {
	incrementAllValuesByOne(grid)

	expectedFlashCount := len(*grid) * len((*grid)[0])

	// Could do something recursive but let's just brute force.

	var flashes int
	for {
		var flashesThisRun int
		for rowIndex := 0; rowIndex < len(*grid); rowIndex++ {
			row := (*grid)[rowIndex]
			for colIndex := 0; colIndex < len(row); colIndex++ {
				val := row[colIndex]
				if val > 9 {
					row[colIndex] = 0
					flashesThisRun++
					flashes++

					incrementNeighbours(grid, rowIndex, colIndex)
				}
			}
		}

		if flashesThisRun == 0 {
			break
		}
	}

	return flashes == expectedFlashCount
}

func printGrid(grid *[][]byte) {
	for rowIndex := 0; rowIndex < len(*grid); rowIndex++ {
		fmt.Println((*grid)[rowIndex])
	}
}

func day11() {
	grid := parseFileToGrid("day11.txt")
	gridPointer := &grid

	printGrid(gridPointer)

	var flashes int
	numDays := 100
	for i := 0; i < numDays; i++ {
		flashes += incrementDay11GridCountingFlashes(gridPointer)

		fmt.Println("===================")
		fmt.Println("Grid at day", i+1)
		printGrid(gridPointer)

		fmt.Println()
	}

	fmt.Println("Result is", flashes, "flashes")
}

func day11p2() {
	grid := parseFileToGrid("day11.txt")
	gridPointer := &grid

	numDays := math.MaxInt32
	for i := 0; i < numDays; i++ {
		haveAllFlashed := incrementDay11GridFindSyncPoint(gridPointer)

		if haveAllFlashed {
			fmt.Println("Sync flash happened at step", i+1)
			break
		}
	}
}
