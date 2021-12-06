package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const day4GridSize = 5

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
	val, err := strconv.Atoi(str)
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

func main() {
	day4()
}
