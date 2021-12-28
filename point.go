package main

import "fmt"

type point struct {
	x int
	y int
}

func newPoint(x int, y int) *point {
	return &point{
		x: x,
		y: y,
	}
}

func (val point) String() string {
	return fmt.Sprintf("(%d, %d)", val.x, val.y)
}
