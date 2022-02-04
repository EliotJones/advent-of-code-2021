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

		explodeDay18Recursive(current, 1)

		if current.String() != expected[i] {
			t.Errorf("Expected %s but got %s", expected[i], current.String())
		}
	}
}

func TestDay18WholeSum(t *testing.T) {
	left := parseDay18String("[[[[4,3],4],4],[7,[[8,4],9]]]")
	right := parseDay18String("[1,1]")

	input := addDay18(left, right)

	reduced := reduceDay18(input)

	expected := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"

	if reduced.String() != expected {
		t.Errorf("Expected %s but got %s", expected, reduced.String())
	}
}

func TestCompleteRunDay18(t *testing.T) {
	inputs := [][]string{
		{
			"[1,1]",
			"[2,2]",
			"[3,3]",
			"[4,4]",
		},
		{
			"[1,1]",
			"[2,2]",
			"[3,3]",
			"[4,4]",
			"[5,5]",
		},
		{
			"[1,1]",
			"[2,2]",
			"[3,3]",
			"[4,4]",
			"[5,5]",
			"[6,6]",
		},
		{
			"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
			"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
			"[7,[5,[[3,8],[1,4]]]]",
			"[[2,[2,2]],[8,[8,1]]]",
			"[2,9]",
			"[1,[[[9,3],9],[[9,0],[0,7]]]]",
			"[[[5,[7,4]],7],1]",
			"[[[[4,2],2],6],[8,7]]",
		},
	}

	expecteds := []string{
		"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
	}

	for i, input := range inputs {
		fmt.Println("Day 18 for complete input: ", input)

		result := runToEndDay18(input)
		resultStr := result.String()

		if resultStr != expecteds[i] {
			t.Errorf("Expected %s but got %s", expecteds[i], resultStr)
		}
	}
}

func TestDay18Magnitude(t *testing.T) {
	inputs := []string{
		"[[1,2],[[3,4],5]]",
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
	}

	expected := []int{
		143,
		1384,
		445,
		791,
		1137,
		3488,
	}

	for i, input := range inputs {
		fmt.Println("Day 18 for magnitude input: ", input)

		current := parseDay18String(input)

		result := magnitudeDay18(current)

		if result != expected[i] {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}

func TestSampleInputPart1Day18(t *testing.T) {
	input := []string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}

	expected := 4140

	final := runToEndDay18(input)

	result := magnitudeDay18(final)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestSampleInputPart2Day18(t *testing.T) {
	input := []string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}

	expected := 3993

	result := locateMaxMagnitude(input)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestRealInputPart1Day18(t *testing.T) {
	scanner, err := scannerForFile("inputs/day18.txt")
	if err != nil {
		panic(err)
	}

	var inputs []string
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	tree := runToEndDay18(inputs)

	result := magnitudeDay18(tree)

	fmt.Println("Result day 18 part 1 is", result)
}

func TestRealInputPart2Day18(t *testing.T) {
	scanner, err := scannerForFile("inputs/day18.txt")
	if err != nil {
		panic(err)
	}

	var inputs []string
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	result := locateMaxMagnitude(inputs)

	fmt.Println("Result day 18 part 2 is", result)
}
