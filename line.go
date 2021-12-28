package main

import (
	"fmt"
)

type line struct {
	start point
	end   point
}

func newLine(x1 int, y1 int, x2 int, y2 int) *line {
	return &line{
		start: *newPoint(x1, y1),
		end:   *newPoint(x2, y2),
	}
}

func (val line) String() string {
	return fmt.Sprintf("(%v) -> (%v)", val.start, val.end)
}

func (val line) PointsOnLine() []point {
	var points []point
	isVertical := val.start.x == val.end.x
	if isVertical {
		start, end := minMax(val.start.y, val.end.y)
		for i := start; i <= end; i++ {
			points = append(points, *newPoint(val.start.x, i))
		}

		return points
	}

	isHorizontal := val.start.y == val.end.y
	if isHorizontal {
		start, end := minMax(val.start.x, val.end.x)
		for i := start; i <= end; i++ {
			points = append(points, *newPoint(i, val.start.y))
		}

		return points
	}

	var dy, dx int
	if (val.end.y - val.start.y) >= 0 {
		dy = 1
	} else {
		dy = -1
	}

	if (val.end.x - val.start.x) >= 0 {
		dx = 1
	} else {
		dx = -1
	}

	// 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9.
	xStart, xEnd := minMax(val.start.x, val.end.x)
	xSteps := xEnd - xStart
	for i := 0; i <= xSteps; i++ {
		p := *newPoint(val.start.x+(i*dx), val.start.y+(i*dy))
		points = append(points, p)
	}

	return points
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
