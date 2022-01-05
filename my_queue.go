package main

import "math"

type distancedElement struct {
	value    interface{}
	distance int
}

type lowestDistanceQueue map[int32]*distancedElement

func (q *lowestDistanceQueue) update(key int32, newDistance int) {
	(*q)[key].distance = newDistance
}

func (q *lowestDistanceQueue) push(key int32, value interface{}, distance int) {
	(*q)[key] = &distancedElement{
		value:    value,
		distance: distance,
	}
}

func (q *lowestDistanceQueue) pop() (bool, *distancedElement) {
	var min *distancedElement
	minDistance, minKey := math.MaxInt32, int32(-1)
	for k, v := range *q {
		if v == nil {
			continue
		}

		if v.distance < minDistance {
			minDistance = v.distance
			min = v
			minKey = k
		}
	}

	var hasResult bool
	if min != nil {
		hasResult = true
		(*q)[minKey] = nil
	}

	return hasResult, min
}
