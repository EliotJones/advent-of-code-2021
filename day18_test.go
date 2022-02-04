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
		fmt.Println("Day 18 for parse input: ", expected)

		tree := parseDay18String(expected)

		output := tree.String()

		if expected != output {
			t.Errorf("Expected output of %s but was %s", expected, output)
		}
	}
}

func TestDay18Add(t *testing.T) {
	left := []string{
		"[1,2]",
		"[[[[4,3],4],4],[7,[[8,4],9]]]",
	}

	right := []string{
		"[[3,4],5]",
		"[1,1]",
	}

	expected := []string{
		"[[1,2],[[3,4],5]]",
		"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
	}

	for i, l := range left {
		fmt.Println("Day 18 for add input: ", l)

		leftRoot := parseDay18String(l)
		rightRoot := parseDay18String(right[i])

		result := addDay18(leftRoot, rightRoot)

		output := result.String()

		if output != expected[i] {
			t.Errorf("Expected %s but got %s", expected[i], output)
		}
	}
}

func TestDay18Reduce(t *testing.T) {
	inputs := []string{
		"[[[[[9,8],1],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]",
		"[[6,[5,[4,[3,2]]]],1]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
	}

	expected := []string{
		"[[[[0,9],2],3],4]",
		"[7,[6,[5,[7,0]]]]",
		"[[6,[5,[7,0]]],3]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}

	for i, input := range inputs {
		fmt.Println("Day 18 for reduce input: ", input)

		current := parseDay18String(input)

		reduced := reduceDay18(current)

		if reduced.String() != expected[i] {
			t.Errorf("Expected %s but got %s", expected[i], reduced.String())
		}
	}
}
