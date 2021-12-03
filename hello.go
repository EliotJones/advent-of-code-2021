package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println("Failed to open the file", err)
		os.Exit(1)
	}

	changeCount := 0
	prev := -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if prev >= 0 && val > prev {
			changeCount++
		}
		prev = val
	}

	fmt.Println(changeCount)
}
