package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	day3()
}
