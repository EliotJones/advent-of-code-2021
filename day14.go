package main

import (
	"fmt"
	"math"
	"strings"
)

type day14Rule struct {
	from1 byte
	from2 byte
	to    byte
}

func parseDay14Input(path string) (string, []day14Rule) {
	scanner, err := scannerForFile(path)
	if err != nil {
		panic(err)
	}

	var input string
	var rules []day14Rule

	var inRules bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inRules = true
			continue
		}

		if inRules {
			parts := strings.Split(line, "->")

			from := strings.Trim(parts[0], " ")

			rules = append(rules, *&day14Rule{
				from1: from[0],
				from2: from[1],
				to:    strings.Trim(parts[1], " ")[0],
			})
		} else {
			input = strings.Trim(line, " ")
		}
	}

	return input, rules
}

func insert(a []byte, index int, value byte) []byte {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func day14MinMaxCommon(buff []byte) (int, int) {
	counts := make(map[byte]int, 4)
	for _, b := range buff {
		if _, ok := counts[b]; ok {
			counts[b]++
		} else {
			counts[b] = 1
		}
	}

	min, max := math.MaxInt32, 0
	for _, v := range counts {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return min, max
}

func day14() {
	input, rules := parseDay14Input("inputs/day14.txt")

	fmt.Println("input", input)
	fmt.Println("rules", rules)

	buff := []byte(input)
	iterations := 10
	for i := 0; i < iterations; i++ {
		fmt.Println("Iteration number", i+1)

		var offset int
		fixedLength := len(buff)
		for j := 0; j < fixedLength-1; j++ {
			p1, p2 := buff[j+offset], buff[j+offset+1]

			var ins byte
			for _, p := range rules {
				if p.from1 == p1 && p.from2 == p2 {
					ins = p.to
				}
			}

			buff = insert(buff, j+offset+1, ins)

			offset++
		}

		fmt.Println("After iteration", i+1, "length is", len(buff))
	}

	min, max := day14MinMaxCommon(buff)

	fmt.Println("Min", min, "Max", max)
	fmt.Println("Result is", max-min)
}
