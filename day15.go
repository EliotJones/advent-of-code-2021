package main

import (
	"fmt"
)

type highScore struct {
	val int
}

func parseDay15Input(path string) [][]byte {
	scanner, err := scannerForFile(path)
	if err != nil {
		panic(err)
	}

	var result [][]byte
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		for i := 0; i < len(line); i++ {
			line[i] -= asciiNumToByteAdjustment
		}

		result = append(result, line)
	}

	return result
}

func getTopRightEdgeAndBottomLeftEdgeAndSimpleMinScore(grid [][]byte) (int, int, int) {
	var topRight, bottomLeft, leastScoringNaive int
	height := len(grid)

	// Top right.
	row := grid[0]

	for i, v := range row {
		if i != 0 {
			topRight += int(v)
		}
	}

	for i, r := range grid {
		if i != 0 {
			topRight += int(r[len(r)-1])
			bottomLeft += int(r[0])
		}
	}

	// Bottom left.
	row = grid[height-1]
	for i, v := range row {
		if i != 0 {
			bottomLeft += int(v)
		}
	}

	x, y := 0, 0
	for {
		nextX, nextY := byte(255), byte(255)
		if x < len(row)-1 {
			nextX = grid[y][x+1]
		}
		if y < height-1 {
			nextY = grid[y+1][x]
		}

		if nextX <= nextY {
			leastScoringNaive += int(nextX)
			x++
		} else {
			leastScoringNaive += int(nextY)
			y++
		}

		if x == len(row)-1 && y == height-1 {
			break
		}
	}

	return topRight, bottomLeft, leastScoringNaive
}

func tooHighScoring(score int, results *[]int) bool {
	for _, r := range *results {
		if r < score {
			return true
		}
	}

	return false
}

func keyForLocation(x int, y int) int16 {
	return (int16(x) << 8) + int16(y)
}

func moveNext(x int, y int, grid *[][]byte, score int, width int, height int, result *highScore, visited map[int16]bool) {
	score += int((*grid)[x][y])

	if score > (*result).val {
		return
	}

	// At end
	if y == height-1 && x == width-1 {
		fmt.Println("Reached end with new score of", score)
		(*result).val = score
		return
	}

	distance := ((height - 1) - y) + ((width - 1) - x) + 1
	if score+distance > (*result).val {
		return
	}

	key := keyForLocation(x, y)
	if v, ok := visited[key]; ok {
		if v {
			// Do not backtrack.
			// fmt.Println("Backtracked, exiting")
			return
		}
	}

	visited[key] = true

	// // Copy visited map
	// newVisited := make(map[int16]struct{}, len(visited)+1)
	// for k, v := range visited {
	// 	newVisited[k] = v
	// }

	// newVisited[key] = empty

	// if x > 0 {
	// 	moveNext(x-1, y, grid, score, results, newVisited)
	// }
	if x < width-1 {
		moveNext(x+1, y, grid, score, width, height, result, visited)
	}
	// if y > 0 {
	// 	moveNext(x, y-1, grid, score, results, newVisited)
	// }

	if y < height-1 {
		moveNext(x, y+1, grid, score, width, height, result, visited)
	}

	visited[key] = false
}

func day15() {
	grid := parseDay15Input("inputs/day15-0.txt")

	topRight, bottomLeft, simpleMin := getTopRightEdgeAndBottomLeftEdgeAndSimpleMinScore(grid)

	fmt.Println("Scores to beat (tr, bl, sim)", topRight, bottomLeft, simpleMin)

	min, _ := minMax(topRight, bottomLeft)
	min, _ = minMax(min, simpleMin)

	// We have an upper bound on the score, now let's explore the valid search space recursively.
	fmt.Println("Min score so far", min)

	// Ignore start position
	score := -1 * int(grid[0][0])
	visited := make(map[int16]bool, 0)
	result := &highScore{
		val: min,
	}

	moveNext(0, 0, &grid, score, len(grid[0]), len(grid), result, visited)

	fmt.Println("Result is", result.val)
}
