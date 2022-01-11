package main

import (
	"fmt"
	"strings"
)

func parseDay17Input(value string) (point, point) {
	xIndex := strings.Index(value, "x=")
	yIndex := strings.Index(value, "y=")

	xParts := value[xIndex+2 : yIndex-2]
	yParts := value[yIndex+2:]

	xs := strings.Split(xParts, "..")
	ys := strings.Split(yParts, "..")

	p1 := *newPoint(parseInt(xs[0]), parseInt(ys[0]))
	p2 := *newPoint(parseInt(xs[1]), parseInt(ys[1]))

	return p1, p2
}

func getHighestPointAndDoesInterceptTarget(x int, y int, t1 point, t2 point) (point, bool) {
	var maxY, xPos, yPos, maxYXPos int
	step := 0

	var hasEncounteredTarget bool

	minTx, maxTx := minMax(t1.x, t2.x)
	minTy, maxTy := minMax(t1.y, t2.y)

	for {
		xPos += x
		yPos += y

		if yPos > maxY {
			maxY = yPos
			maxYXPos = xPos
		}

		if x > 0 {
			x--
		} else if x < 0 {
			x++
		}

		y--

		if xPos > maxTx {
			break
		}

		if yPos < minTy {
			break
		}

		if xPos >= minTx && xPos <= maxTx && yPos >= minTy && yPos <= maxTy {
			hasEncounteredTarget = true
		}

		step++
	}

	return *newPoint(maxYXPos, maxY), hasEncounteredTarget
}

func day17() {
	scanner, err := scannerForFile("inputs/day17.txt")
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	line := scanner.Text()

	p1, p2 := parseDay17Input(line)

	fmt.Println("target area is", p1, p2)

	_, maxTx := minMax(p1.x, p2.x)
	minTy, _ := minMax(p1.y, p2.y)

	// I had to look up some clues for the bounds here, https://work.njae.me.uk/2021/12/19/advent-of-code-2021-day-17/ :(
	maxPossibleX := maxTx
	maxPossibleY := abs(minTy)

	maxY := 0

	for y := 1; y < maxPossibleY+1; y++ {
		for x := 1; x < maxPossibleX+1; x++ {

			max, inTarget := getHighestPointAndDoesInterceptTarget(x, y, p1, p2)

			if !inTarget {
				continue
			}

			if max.y > maxY {
				maxY = max.y
			}
		}
	}

	fmt.Println("Encountered", maxY)
}
