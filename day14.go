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

type day14InsertionPair struct {
	left  byte
	right byte
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

func day14MinMaxCommonPart2(pairs map[int16]int64, start byte, end byte) (int64, int64) {
	counts := make(map[byte]int64, 4)
	for k, v := range pairs {
		b1, b2 := byte(k>>8), byte(k)

		if _, ok := counts[b1]; ok {
			counts[b1] += v
		} else {
			counts[b1] = v
		}

		if _, ok := counts[b2]; ok {
			counts[b2] += v
		} else {
			counts[b2] = v
		}
	}

	var min, max int64
	min, max = math.MaxInt64, 0
	for b, v := range counts {
		if b == start || b == end {
			v -= 1
		}

		char := string([]byte{b})
		fmt.Print(char)
		actual := v / 2

		if b == start || b == end {
			actual += 1
		}
		if actual > max {
			max = actual
		}

		if actual < min {
			min = actual
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

func inputToPairMap(input string) map[int16]int64 {
	result := make(map[int16]int64, len(input)-1)
	for i := 0; i < len(input)-1; i++ {
		key := pairToKey(input[i], input[i+1])
		if _, ok := result[key]; ok {
			result[key]++
		} else {
			result[key] = 1
		}
	}

	return result
}

func rulesToMap(rules []day14Rule) map[int16]byte {
	m := make(map[int16]byte, len(rules))

	for _, r := range rules {
		k := pairToKey(r.from1, r.from2)
		m[k] = r.to
	}

	return m
}

func pairToKey(b1 byte, b2 byte) int16 {
	return int16(b1)<<8 + int16(b2)
}

func printDay14Map(m map[int16]int64) {
	for k, v := range m {
		if v == 0 {
			continue
		}

		str := string([]byte{byte(k >> 8), byte(k)})
		fmt.Print(str, ": ", v, ", ")
	}
}

func day14p2() {
	input, rules := parseDay14Input("inputs/day14.txt")

	fmt.Println("input", input)
	fmt.Println("rules", rules)

	pairCountMap := inputToPairMap(input)
	rulesMap := rulesToMap(rules)

	start, end := input[0], input[len(input)-1]

	iterations := 40
	for i := 0; i < iterations; i++ {
		insertPairs := make(map[int16]int64, 0)
		for k, v := range pairCountMap {
			if v == 0 {
				continue
			}

			b1, b2 := byte(k>>8), byte(k)
			newByte := rulesMap[k]
			newKey1, newKey2 := (int16(b1)<<8)+int16(newByte), (int16(newByte)<<8)+int16(b2)

			if _, ok := insertPairs[newKey1]; ok {
				insertPairs[newKey1] += v
			} else {
				insertPairs[newKey1] = v
			}

			if _, ok := insertPairs[newKey2]; ok {
				insertPairs[newKey2] += v
			} else {
				insertPairs[newKey2] = v
			}

			pairCountMap[k] -= v
		}

		for k, v := range insertPairs {
			if _, ok := pairCountMap[k]; ok {
				pairCountMap[k] += v
			} else {
				pairCountMap[k] = v
			}
		}

		fmt.Println("Finished iteration", i+1, "up to", len(pairCountMap), "pairs")
	}

	min, max := day14MinMaxCommonPart2(pairCountMap, start, end)
	fmt.Println("Result is", max-min)
}
