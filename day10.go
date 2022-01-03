package main

import (
	"fmt"
	"sort"
)

type openChunk struct {
	open       string
	startIndex int
	close      string
}

type chunkStack []openChunk

func (s chunkStack) Push(val openChunk) chunkStack {
	return append(s, val)
}

func (s chunkStack) Pop() (chunkStack, openChunk) {
	length := len(s)
	if length == 0 {
		return s, openChunk{}
	}

	return s[:length-1], s[length-1]
}

func (s chunkStack) Peek() openChunk {
	length := len(s)
	if length == 0 {
		return openChunk{}
	}

	return s[length-1]
}

func containsStr(v string, m map[string]string) bool {
	if _, ok := m[v]; ok {
		return true
	}

	return false
}

func parseLineReturnFirstInvalidOrStack(line string, openClose map[string]string) (string, chunkStack) {
	stack := make(chunkStack, 0)

	for i := 0; i < len(line); i++ {
		c := line[i : i+1]
		if value, ok := openClose[c]; ok {
			stack = stack.Push(openChunk{
				open:       c,
				startIndex: i,
				close:      value,
			})
		} else if len(stack) > 0 {
			top := stack.Peek()
			if top.close == c {
				stack, _ = stack.Pop()
			} else {
				return c, stack
			}
		} else {
			return c, stack
		}
	}

	return "", stack
}

func getOpenCloseMapping() map[string]string {
	openClose := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	return openClose
}

func day10() {
	openClose := getOpenCloseMapping()

	charScores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	scanner, err := scannerForFile("inputs/day10.txt")
	if err != nil {
		panic(err)
	}

	var result int
	for scanner.Scan() {
		invalid, _ := parseLineReturnFirstInvalidOrStack(scanner.Text(), openClose)

		if invalid == "" {
			continue
		}

		result += charScores[invalid]
	}

	fmt.Println("Result is", result)
}

func day10p2() {
	openClose := getOpenCloseMapping()

	charScores := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	scanner, err := scannerForFile("inputs/day10.txt")
	if err != nil {
		panic(err)
	}

	var scores []int64
	for scanner.Scan() {
		invalid, stack := parseLineReturnFirstInvalidOrStack(scanner.Text(), openClose)

		if invalid != "" {
			continue
		}

		var result int
		for {
			if len(stack) > 0 {
				var top openChunk
				stack, top = stack.Pop()

				result *= 5
				result += charScores[top.close]
			} else {
				scores = append(scores, int64(result))
				break
			}
		}
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })

	mid := scores[len(scores)/2]

	fmt.Println("Result is", mid)
}
