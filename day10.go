package main

import "fmt"

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

func parseLineReturnFirstInvalid(line string, openClose map[string]string) string {
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
				return c
			}
		} else {
			return c
		}
	}

	return ""
}

func day10() {
	charScores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	openClose := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	scanner, err := scannerForFile("day10.txt")
	if err != nil {
		panic(err)
	}

	var result int
	for scanner.Scan() {
		invalid := parseLineReturnFirstInvalid(scanner.Text(), openClose)

		if invalid == "" {
			continue
		}

		result += charScores[invalid]
	}

	fmt.Println("Result is", result)
}
