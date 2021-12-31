package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const day4GridSize = 5
const day8OneLength = 2
const day8FourLength = 4
const day8SevenLength = 3
const day8EightLength = 7

type binaryNode struct {
	children   [2]*binaryNode
	value      int
	childCount int
	isRoot     bool
}

type trie struct {
	root *binaryNode
}

type bingoIndexEntry struct {
	board int
	index int
}

type day8InputOutput struct {
	input  []string
	output []string
}

func (t trie) insert(value []byte, correction int) {
	current := t.root
	for i := 0; i < len(value); i++ {
		index := int(value[i]) + correction
		if current.children[index] == nil {
			current.children[index] = &binaryNode{
				value: index,
			}
		}

		current = current.children[index]
		current.childCount++
	}
}

func scannerForFile(path string) (*bufio.Scanner, error) {
	if file, err := os.Open(path); err == nil {
		return bufio.NewScanner(file), nil
	} else {
		return nil, err
	}
}

func day1() {
	scanner, err := scannerForFile("day1.txt")
	if err != nil {
		fmt.Println("Failed to scan file", err)
		os.Exit(1)
	}

	var index = 0
	var preceding [3]int
	var windowSum, prevSum int
	changeCount := 0
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		if index < 3 {
			preceding[index] = val
			windowSum += val
		} else {
			replacementIndex := index % 3
			prevSum = windowSum
			windowSum = windowSum - preceding[replacementIndex] + val
			preceding[replacementIndex] = val

			if windowSum > prevSum {
				changeCount++
			}
		}

		index++
	}

	fmt.Println(changeCount)
}

func day2() {
	scanner, err := scannerForFile("day2.txt")
	if err != nil {
		fmt.Println(err)
	}

	var length, depth, aim int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			fmt.Println("Bad input " + line)
			os.Exit(1)
		}

		if val, err := strconv.Atoi(parts[1]); err == nil {
			switch parts[0] {
			case "down":
				aim += val
				break
			case "forward":
				length += val
				depth += aim * val
				break
			case "up":
				aim -= val
				break
			}
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println("Depth is ", depth, " length is ", length, ". Total is ", depth*length)
}

func day3() {
	scanner, err := scannerForFile("day3.txt")
	if err != nil {
		panic(err)
	}

	const zeroValue = 48

	var sums []int
	var count int
	hasInit := false
	for scanner.Scan() {
		byteStr := scanner.Bytes()

		for index, b := range byteStr {
			binValue := int(b - zeroValue)
			if !hasInit {
				sums = append(sums, binValue)
			} else {
				sums[index] += binValue
			}
		}
		count++

		hasInit = true
	}

	var gamma, epsilon int
	halfPoint := count / 2

	fmt.Println("Content is", sums, "half point", halfPoint)

	for index, i := range sums {
		var value int
		if i >= halfPoint {
			value = 1
		}

		shift := (len(sums) - index - 1)

		gamma += value << shift
		epsilon += (1 - value) << shift
	}

	fmt.Println(gamma, epsilon, gamma*epsilon)
}

func chooseCorrectNode(node *binaryNode, isOxygen bool) *binaryNode {
	var zeroCount, oneCount int
	if node.children[0] != nil {
		zeroCount = node.children[0].childCount
	}

	if node.children[1] != nil {
		oneCount = node.children[1].childCount
	}

	if zeroCount == 0 && oneCount == 0 {
		return nil
	}

	if node.children[0] == nil {
		return node.children[1]
	} else if node.children[1] == nil {
		return node.children[0]
	}

	var chosenNode *binaryNode
	if zeroCount == oneCount {
		if isOxygen {
			chosenNode = node.children[1]
		} else {
			chosenNode = node.children[0]
		}
	} else if zeroCount > oneCount {
		if isOxygen {
			chosenNode = node.children[0]
		} else {
			chosenNode = node.children[1]
		}
	} else {
		if isOxygen {
			chosenNode = node.children[1]
		} else {
			chosenNode = node.children[0]
		}
	}

	return chosenNode
}

func walkReadingType(node *binaryNode, isOxygen bool, result string) (string, error) {
	if node == nil {
		return result, nil
	}

	bestNode := chooseCorrectNode(node, isOxygen)

	if bestNode == nil {
		return result + strconv.Itoa(node.value), nil
	}

	if !node.isRoot {
		result += strconv.Itoa(node.value)
	}

	return walkReadingType(bestNode, isOxygen, result)
}

func binStrToDec(str string) int {
	var result int
	for i, c := range str {
		val, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		result += val << (len(str) - i - 1)
	}

	return result
}

func day3p2() {
	scanner, err := scannerForFile("day3.txt")
	if err != nil {
		panic(err)
	}

	trie := &trie{
		root: &binaryNode{
			isRoot: true,
		},
	}
	for scanner.Scan() {
		buff := scanner.Bytes()
		trie.insert(buff, -48)
	}

	oxStr, oxErr := walkReadingType(trie.root, true, "")
	if oxErr != nil {
		panic(oxErr)
	}

	co2Str, co2Err := walkReadingType(trie.root, false, "")
	if co2Err != nil {
		panic(co2Err)
	}

	ox, co2 := binStrToDec(oxStr), binStrToDec(co2Str)

	fmt.Println(oxStr, co2Str, ox, co2)

	fmt.Println("result", ox*co2)
}

func parseDay4Input(scanner *bufio.Scanner) ([]string, [][day4GridSize * day4GridSize]string, map[string][]bingoIndexEntry) {
	var boards [][day4GridSize * day4GridSize]string
	var announcedValues []string

	boardIndex, withinBoardIndex := -1, 0
	lookup := make(map[string][]bingoIndexEntry)

	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			announcedValues = strings.Split(line, ",")
			firstLine = false
			continue
		}

		if line == "" {
			boardIndex++
			withinBoardIndex = 0
		} else {
			lineParts := strings.Split(line, " ")

			if len(boards) <= boardIndex {
				var next [day4GridSize * day4GridSize]string
				boards = append(boards, next)
			}

			for _, val := range lineParts {
				if len(val) == 0 {
					continue
				}

				arr, found := lookup[val]
				if !found {
					arr = make([]bingoIndexEntry, 0)
					lookup[val] = arr
				}

				ix := bingoIndexEntry{
					board: boardIndex,
					index: withinBoardIndex,
				}

				lookup[val] = append(arr, ix)

				boards[boardIndex][withinBoardIndex] = val
				withinBoardIndex++
			}
		}
	}

	return announcedValues, boards, lookup
}

func parseInt(str string) int {
	val, err := strconv.Atoi(strings.Trim(str, " "))
	if err != nil {
		return 0
	}

	return val
}

func day4() {

	scanner, err := scannerForFile("day4.txt")

	if err != nil {
		panic(err)
	}

	announcedValues, boards, valuesIndex := parseDay4Input(scanner)

	marked := make([][day4GridSize * day4GridSize]bool, len(boards))

	wonBoards := make([]bool, len(boards))

	var last int
	for callIndex, val := range announcedValues {
		index, found := valuesIndex[val]
		if !found {
			continue
		}

		complete := false
		for _, ix := range index {
			marked[ix.board][ix.index] = true

			if callIndex >= 3 {
				rowStartIndex := ix.index / day4GridSize
				col := ix.index % day4GridSize
				isColComplete, isRowComplete := true, true
				for colIndex := 0; colIndex < day4GridSize; colIndex++ {
					if !marked[ix.board][col+(colIndex*day4GridSize)] {
						isColComplete = false
						break
					}
				}

				for rowIndex := 0; rowIndex < day4GridSize; rowIndex++ {
					if !marked[ix.board][(rowIndex)+(rowStartIndex*day4GridSize)] {
						isRowComplete = false
						break
					}
				}

				if isRowComplete || isColComplete {
					markBoard := marked[ix.board]
					valBoard := boards[ix.board]

					var sum int

					for i := 0; i < day4GridSize*day4GridSize; i++ {
						if !markBoard[i] {
							sum += parseInt(valBoard[i])
						}
					}

					result := sum * parseInt(val)
					last = result
					wonBoards[ix.board] = true

					allWon := true
					for i := 0; i < len(wonBoards); i++ {
						if !wonBoards[i] {
							allWon = false
							break
						}
					}

					if allWon {
						complete = true
						break
					}
				}
			}
		}

		if complete {
			break
		}
	}

	fmt.Println(last)
}

func parseDay5InputLine(input string) (line, error) {
	arrowIndex := strings.Index(input, ">")

	if arrowIndex < 0 {
		return line{}, errors.New("Line did not contain arrow: " + input)
	}

	first := strings.Split(input[0:arrowIndex-2], ",")
	second := strings.Split(input[arrowIndex+1:], ",")

	if len(first) != 2 || len(second) != 2 {
		return line{}, errors.New("Line did not contain 4 points split by an error: " + input)
	}

	x1 := parseInt(first[0])
	y1 := parseInt(first[1])
	x2 := parseInt(second[0])
	y2 := parseInt(second[1])

	return *newLine(x1, y1, x2, y2), nil
}

func day5() {
	scanner, err := scannerForFile("day5.txt")
	if err != nil {
		panic(err)
	}

	pointCountMap := make(map[point]int)

	var result int

	for scanner.Scan() {
		textLine := scanner.Text()

		line, err := parseDay5InputLine(textLine)

		if err != nil {
			panic(err)
		}

		for _, p := range line.PointsOnLine() {
			if val, ok := pointCountMap[p]; ok {
				pointCountMap[p]++
				if val == 1 {
					result++
				}
			} else {
				pointCountMap[p] = 1
			}
		}
	}

	fmt.Println("result is", result)
}

func getIntList(fileName string) ([]int, error) {
	scanner, err := scannerForFile(fileName)
	if err != nil {
		return nil, err
	}

	var result []int
	for scanner.Scan() {
		str := scanner.Text()

		parts := strings.Split(str, ",")

		for _, v := range parts {
			result = append(result, parseInt(v))
		}
	}

	return result, nil
}

func day6ZeroAge(day int) [7]int64 {
	var ageCounts [9]int
	var last7Counts [7]int64

	// One fish of age zero to begin
	ageCounts[0] = 1

	for i := 0; i < day; i++ {

		// Increment each counter, do this in reverse
		var prev int
		var sum int64
		for j := 8; j >= 0; j-- {
			count := ageCounts[j]

			if j < 8 {
				ageCounts[j] = prev
				sum += int64(prev)
				prev = 0
			} else {
				ageCounts[j] = 0
			}

			prev = count

			if j == 0 {
				ageCounts[8] += prev
				ageCounts[6] += prev

				sum += int64(prev * 2)
			}
		}

		if i >= day-7 {
			last7Counts[7-(day-i)] = sum
		}
	}

	return last7Counts
}

func day6Simple(day int) {
	scanner, err := scannerForFile("day6.txt")
	if err != nil {
		panic(err)
	}

	agesCountMap := make(map[int]int)

	for scanner.Scan() {
		text := scanner.Text()

		parts := strings.Split(text, ",")
		for _, v := range parts {
			age := parseInt(v)
			if cnt, ok := agesCountMap[age]; ok {
				agesCountMap[age] = cnt + 1
			} else {
				agesCountMap[age] = 1
			}
		}
	}

	// Get how many of each starting age
	fmt.Println(agesCountMap)

	// Calculate the number of fish for a single fish of starting age 0 on day n (and counts for n - 6 days prior)
	last7DayCounts := day6ZeroAge(day)

	// Adjust each input age by the correct index from the last 7 days (for age 0 fish) and number with that age.
	var result int64
	for k, v := range agesCountMap {
		countOnDay := last7DayCounts[6-k]
		result += int64(v) * countOnDay
	}

	fmt.Println("Result for day", day, "is", result)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func getMinMaxFromSlice(slice []int) (int, int) {
	lowest := math.MaxInt32
	var highest int
	for i := 0; i < len(slice); i++ {
		val := slice[i]
		if val < lowest {
			lowest = val
		}

		if val > highest {
			highest = val
		}
	}

	return lowest, highest
}

func day7() {
	values, err := getIntList("day7.txt")
	if err != nil {
		panic(err)
	}

	min, max := getMinMaxFromSlice(values)

	lowestSum := math.MaxInt32
	index := 0
	for i := min; i <= max; i++ {
		var sumForMove int
		for j := 0; j < len(values); j++ {
			value := values[j]
			if i == value {
				continue
			}

			gap := abs(value - i)
			sumForMove += gap
		}

		if sumForMove < lowestSum {
			lowestSum = sumForMove
			index = i
		}
	}

	fmt.Println("lowest sum is", lowestSum, "at index", index)
}

func day7p2() {
	values, err := getIntList("day7.txt")
	if err != nil {
		panic(err)
	}

	min, max := getMinMaxFromSlice(values)

	lowestSum := math.MaxInt32
	index := 0
	for i := min; i <= max; i++ {
		var sumForMove int
		for j := 0; j < len(values); j++ {
			value := values[j]
			if i == value {
				continue
			}

			gap := abs(value-i) + 1
			sumForMove += (gap * (gap - 1)) / 2
		}

		if sumForMove < lowestSum {
			lowestSum = sumForMove
			index = i
		}
	}

	fmt.Println("lowest sum is", lowestSum, "at index", index)
}

func parseDay8InputLines(scanner *bufio.Scanner) []day8InputOutput {
	var result []day8InputOutput
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		input := strings.Split(strings.Trim(parts[0], " "), " ")
		output := strings.Split(strings.Trim(parts[1], " "), " ")
		result = append(result, day8InputOutput{
			input:  input,
			output: output,
		})
	}

	return result
}

func day8() {

	scanner, err := scannerForFile("day8.txt")
	if err != nil {
		panic(err)
	}

	lines := parseDay8InputLines(scanner)

	var result int
	for _, line := range lines {
		for _, str := range line.output {
			length := len(str)
			if length == day8OneLength || length == day8FourLength || length == day8SevenLength || length == day8EightLength {
				result++
			}
		}
	}

	fmt.Println("Result is", result)
}

func day8p2() {
	// Digital display numbers as binary,
	// values are in order top, topLeft, topRight, middle, bottomLeft, bottomRight, bottom.
	digitalDisplayNumbersByteEncoded := [10]byte{
		byte(0b1110111),
		byte(0b0010010),
		byte(0b1011101),
		byte(0b1011011),
		byte(0b0111010),
		byte(0b1101011),
		byte(0b1101111),
		byte(0b1010010),
		byte(0b1111111),
		byte(0b1111011),
	}

	scanner, err := scannerForFile("day8.txt")
	if err != nil {
		panic(err)
	}

	lines := parseDay8InputLines(scanner)

	// Get known values (1, 4, 7, 8) and values by number of segments from input.
	var overallResult int
	for _, line := range lines {
		var one, two, three, four, five, seven, eight map[byte]struct{}
		var fiveLengths, sixLengths []map[byte]struct{}
		for _, str := range line.input {
			if len(str) == day8OneLength {
				one = stringToMap(str)
			} else if len(str) == day8FourLength {
				four = stringToMap(str)
			} else if len(str) == day8SevenLength {
				seven = stringToMap(str)
			} else if len(str) == day8EightLength {
				eight = stringToMap(str)
			} else if len(str) == 6 {
				sixLengths = append(sixLengths, stringToMap(str))
			} else if len(str) == 5 {
				fiveLengths = append(fiveLengths, stringToMap(str))
			}
		}

		// Start finding corresponding segments.
		var top, topRight, topLeft, middle, bottomLeft, bottomRight, bottom byte
		var sixIndex int

		// 7 is 1 plus the line at the top.
		top = except(seven, one)[0]
		// 6 is the only digit of length 6 missing a part of 1 (the top right).
		topRight, sixIndex = deduceTopRightAndSixIndex(sixLengths, one)
		// The other non-top-right element of one is the bottom right.
		bottomRight = deduceBottomRight(topRight, one)

		// Length 5 digits are 2, 3 and 5.
		for _, m := range fiveLengths {
			// Of length 5 digits, only 3 completely excludes 1.
			diff := except(one, m)
			if len(diff) == 0 {
				three = m
			} else if contains(m, bottomRight) {
				// 2 is missing bottom right so this must be 5.
				five = m
			} else if contains(m, topRight) {
				// 5 is missing top right so this must be 2.
				two = m
			}
		}

		// Subtract 5 from 2 and you have bottom left and top right.
		bottomLeft = deduceBottomLeft(two, five, topRight)
		bottom = deduceBottom(four, seven, eight, bottomLeft)

		// Length 6 digits are 0, 6 and 9.
		for i, m := range sixLengths {
			if i == sixIndex {
				continue
			}

			// 8 minus 9 leaves bottom left, 8 minus 0 leaves middle.
			diff := except(eight, m)
			if diff[0] == bottomLeft {
				topLeft = except(m, three)[0]
			} else {
				middle = diff[0]
			}
		}

		// All values are now found, map back to binary representation.
		bytesToShiftMap := map[byte]int{
			top:         0,
			topLeft:     1,
			topRight:    2,
			middle:      3,
			bottomLeft:  4,
			bottomRight: 5,
			bottom:      6,
		}

		lineResult := mapDisplayOutputToInt(bytesToShiftMap, line.output, digitalDisplayNumbersByteEncoded)
		overallResult += lineResult
	}

	fmt.Println("Result is", overallResult)
}

func main() {
	day10p2()
}
