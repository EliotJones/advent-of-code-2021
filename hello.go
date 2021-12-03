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

func main() {
	day2()
}
