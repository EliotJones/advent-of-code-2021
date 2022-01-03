package main

import (
	"fmt"
	"strings"
)

type day13Fold struct {
	dimension byte
	location  int
}

func parseDay13Input(path string) ([]point, []day13Fold) {
	scanner, err := scannerForFile(path)

	if err != nil {
		panic(err)
	}

	var folds []day13Fold
	var points []point
	var inFolds bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inFolds = true
			continue
		}

		if inFolds {
			alongIndex := strings.Index(line, "along")
			if alongIndex < 0 {
				continue
			}

			subStr := strings.Trim(line[alongIndex+1+len("along"):], " ")
			parts := strings.Split(subStr, "=")

			var dimension byte
			if parts[0] == "y" {
				dimension = 1
			}

			folds = append(folds, *&day13Fold{
				dimension: dimension,
				location:  parseInt(parts[1]),
			})
		} else {
			parts := strings.Split(line, ",")
			if len(parts) != 2 {
				continue
			}

			x, y := parseInt(parts[0]), parseInt(parts[1])
			points = append(points, *newPoint(x, y))
		}
	}

	return points, folds
}

func getMaxXAndY(points []point) (int, int) {
	maxX, maxY := 0, 0
	for _, p := range points {
		if p.x > maxX {
			maxX = p.x
		}

		if p.y > maxY {
			maxY = p.y
		}
	}

	return maxX, maxY
}

func pointAt(points []point, x int, y int) bool {
	for _, p := range points {
		if x == p.x && y == p.y {
			return true
		}
	}

	return false
}

func printGridDay13(points []point, maxX int, maxY int) {
	for rowIndex := 0; rowIndex <= maxY; rowIndex++ {
		for colIndex := 0; colIndex <= maxX; colIndex++ {
			if pointAt(points, colIndex, rowIndex) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func countValidDay13Points(points []point, removedVal int) int {
	var count int
	for _, p := range points {
		if p.x != removedVal {
			count++
		}
	}

	return count
}

func day13() {
	const removedVal = -1

	points, folds := parseDay13Input("inputs/day13.txt")
	maxX, maxY := getMaxXAndY(points)

	foldsToRun := len(folds)
	for i := 0; i < foldsToRun; i++ {
		fold := folds[i]

		isXFold := fold.dimension == 0

		var newZero bool
		if isXFold && fold.location > maxX/2 {
			newZero = true
		} else if !isXFold && fold.location > maxY/2 {
			newZero = true
		}

		fmt.Println("Fold at", fold, "has new zero val", newZero)

		maxNewIndex := fold.location - 1
		fixedLen := len(points)
		for j := 0; j < fixedLen; j++ {
			point := points[j]
			if point.x == removedVal {
				continue
			}

			var pointLoc int
			if isXFold && point.x > fold.location {
				pointLoc = point.x
			} else if !isXFold && point.y > fold.location {
				pointLoc = point.y
			} else {
				continue
			}

			diff := pointLoc - fold.location

			newLoc := maxNewIndex - (diff - 1)

			if isXFold && !pointAt(points, newLoc, point.y) {
				points[j].x = newLoc
			} else if !isXFold && !pointAt(points, point.x, newLoc) {
				points[j].y = newLoc
			} else {
				points[j].x = removedVal
				points[j].y = removedVal
			}
		}

		maxX, maxY = getMaxXAndY(points)

		fmt.Println("After fold", i, "at", fold.location, "there are", countValidDay13Points(points, removedVal), "points")
	}

	printGridDay13(points, maxX, maxY)
}
