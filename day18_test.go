package main

import (
	"fmt"
	"testing"
)

func TestDay18Parse(t *testing.T) {
	inputs := []string{
		"[1,2]",
		"[[1,2],3]",
		"[9,[8,7]]",
		"[[1,9],[8,5]]",
		"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
		"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
		"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
	}

	for _, expected := range inputs {
		fmt.Println("Day 18 for input: ", expected)

		tree := parseDay18String(expected)

		output := tree.String()

		if expected != output {
			t.Errorf("Expected output of %s but was %s", expected, output)
		}
	}
}
