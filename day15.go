package main

import (
	"fmt"
)

type day15Node struct {
	id     int32
	weight int
	links  []*day15Node
}

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

func parseDay15InputToGraphAndGetEndId(path string, repeats int) (*day15Node, int32) {
	scanner, err := scannerForFile(path)
	if err != nil {
		panic(err)
	}

	lines := make([]string, 0)

	var lastId int32
	nodes := make(map[int32]*day15Node, 0)
	var rowIndex, colIndex int
	for scanner.Scan() {
		// Bufio breaks (expected but bad) if using bytes here: https://stackoverflow.com/questions/24919968/strange-behavior-of-buffo-scanner-reading-file-line-by-line.
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		lines = append(lines, line)
	}

	for r := 0; r < repeats; r++ {
		for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
			line := lines[lineIndex]

			colIndex = 0
			for r2 := 0; r2 < repeats; r2++ {
				for i := 0; i < len(line); i++ {
					b := line[i] - asciiNumToByteAdjustment

					riskAdjustment := byte((r) + (r2))

					if b+riskAdjustment > 9 {
						b = (b + riskAdjustment) - 9
					} else {
						b += riskAdjustment
					}

					id := keyForLocation(colIndex, rowIndex)

					node := day15Node{
						id:     id,
						weight: int(b),
						links:  make([]*day15Node, 0),
					}

					nodePointer := &node

					if colIndex > 0 {
						prevKey := keyForLocation(colIndex-1, rowIndex)
						prevNode := nodes[prevKey]
						node.links = append(node.links, prevNode)
						(*prevNode).links = append((*prevNode).links, nodePointer)
					}

					if rowIndex > 0 {
						prevKey := keyForLocation(colIndex, rowIndex-1)
						prevNode := nodes[prevKey]
						node.links = append(node.links, prevNode)
						(*prevNode).links = append((*prevNode).links, nodePointer)
					}

					nodes[id] = nodePointer
					lastId = id
					colIndex++
				}
			}

			rowIndex++
		}
	}

	// Get the start node
	return nodes[0], lastId
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

func keyForLocation(x int, y int) int32 {
	return (int32(x) << 16) + int32(y)
}

func moveNext(x int, y int, grid *[][]byte, score int, width int, height int, result *highScore, visited map[int32]bool) {
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
			return
		}
	}

	visited[key] = true

	if x > 0 {
		moveNext(x-1, y, grid, score, width, height, result, visited)
	}
	if x < width-1 {
		moveNext(x+1, y, grid, score, width, height, result, visited)
	}
	if y > 0 {
		moveNext(x, y-1, grid, score, width, height, result, visited)
	}
	if y < height-1 {
		moveNext(x, y+1, grid, score, width, height, result, visited)
	}

	visited[key] = false
}

func dijkstra(queue *lowestDistanceQueue, initial *day15Node, endId int32) {
	distances := make(map[int32]int)
	distances[initial.id] = 0

	for {
		hasPop, elem := (*queue).pop()
		if !hasPop {
			break
		}

		node := elem.value.(*day15Node)
		for _, link := range node.links {
			newDistance := link.weight + elem.distance
			if d, ok := distances[link.id]; ok {
				if newDistance < d {

					distances[link.id] = newDistance
					(*queue).push(link.id, link, newDistance)
				}
			} else {
				distances[link.id] = newDistance
				(*queue).push(link.id, link, newDistance)
			}

		}
	}

	res := distances[endId]
	fmt.Println("Result is", res)
}

func day15BruteForce() {
	grid := parseDay15Input("inputs/day15-0.txt")

	topRight, bottomLeft, simpleMin := getTopRightEdgeAndBottomLeftEdgeAndSimpleMinScore(grid)

	fmt.Println("Scores to beat (tr, bl, sim)", topRight, bottomLeft, simpleMin)

	min, _ := minMax(topRight, bottomLeft)
	min, _ = minMax(min, simpleMin)

	// We have an upper bound on the score, now let's explore the valid search space recursively.
	fmt.Println("Min score so far", min)

	// Ignore start position
	score := -1 * int(grid[0][0])
	visited := make(map[int32]bool, 0)
	result := &highScore{
		val: min,
	}

	moveNext(0, 0, &grid, score, len(grid[0]), len(grid), result, visited)

	fmt.Println("Result is", result.val)
}

func day15() {
	startNode, endId := parseDay15InputToGraphAndGetEndId("inputs/day15.txt", 1)

	queue := make(lowestDistanceQueue, 0)

	queue.push(startNode.id, startNode, 0)

	dijkstra(&queue, startNode, endId)
}

func day15p2() {
	startNode, endId := parseDay15InputToGraphAndGetEndId("inputs/day15.txt", 5)

	queue := make(lowestDistanceQueue, 0)

	queue.push(startNode.id, startNode, 0)

	dijkstra(&queue, startNode, endId)
}
