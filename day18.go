package main

import (
	"strconv"
	"strings"
)

type day18Node struct {
	left       *day18Node
	right      *day18Node
	parent     *day18Node
	leftValue  string
	rightValue string
}

func NewDay18Node() *day18Node {
	return &day18Node{}
}

func printToBuilder(node *day18Node, sb *strings.Builder) {
	sb.WriteString("[")

	if node.leftValue != "" {
		sb.WriteString(node.leftValue)
	} else {
		printToBuilder(node.left, sb)
	}

	sb.WriteString(",")

	if node.rightValue != "" {
		sb.WriteString(node.rightValue)
	} else {
		printToBuilder(node.right, sb)
	}

	sb.WriteString("]")
}

func (n *day18Node) String() string {
	var sb strings.Builder

	printToBuilder(n, &sb)

	return sb.String()
}

func parseDay18String(value string) *day18Node {
	root := NewDay18Node()

	stack := make(genericStack, 0)

	firstOpenBrace := true
	for i := 0; i < len(value); i++ {
		s := value[i : i+1]

		if s == "[" {
			if firstOpenBrace {
				firstOpenBrace = false
				stack.Push(root)
			} else {
				stack.Push(NewDay18Node())
			}
		} else if s == "]" {
			_, me := stack.Pop()

			parent := stack.Peek()
			if parent != nil {
				parentNode := parent.(*day18Node)
				currentNode := me.(*day18Node)

				if parentNode.left == nil && parentNode.leftValue == "" {
					parentNode.left = currentNode
				} else if parentNode.right == nil && parentNode.rightValue == "" {
					parentNode.right = currentNode
				}

				currentNode.parent = parentNode
			}

			continue
		} else if s == "," {
			continue
		} else if s == " " {
			continue
		} else {
			currentTop := stack.Peek().(*day18Node)
			if currentTop.leftValue == "" && currentTop.left == nil {
				currentTop.leftValue = s
			} else if currentTop.right == nil {
				currentTop.rightValue = s
			}
		}
	}

	return root
}

func addDay18(left *day18Node, right *day18Node) *day18Node {
	newRoot := NewDay18Node()
	newRoot.left = left
	newRoot.right = right

	return newRoot
}

func isDay18ValueTooBig(value string) bool {
	intVal := parseInt(value)

	return intVal >= 10
}

func splitDay18Value(value string) *day18Node {
	intVal := parseInt(value)

	left := intVal / 2
	right := intVal - left

	node := NewDay18Node()
	node.leftValue = strconv.Itoa(left)
	node.rightValue = strconv.Itoa(right)

	return node
}

func addDay18Values(one string, two string) string {
	return strconv.Itoa(parseInt(one) + parseInt(two))
}

func findAndAddToFirstInDirection(sourceValue string, current *day18Node, former *day18Node, goLeft bool, descending bool) bool {
	if descending {
		// Swap search direction
		if goLeft && current.right != nil {
			return findAndAddToFirstInDirection(sourceValue, current.right, current, goLeft, descending)
		} else if goLeft && current.rightValue != "" {
			current.rightValue = addDay18Values(sourceValue, current.rightValue)
			return true
		} else if !goLeft && current.left != nil {
			return findAndAddToFirstInDirection(sourceValue, current.left, current, goLeft, descending)
		} else if !goLeft && current.leftValue != "" {
			current.leftValue = addDay18Values(sourceValue, current.leftValue)
			return true
		}
		return false
	}

	if !goLeft && current.right != former && current.right != nil {
		return findAndAddToFirstInDirection(sourceValue, current.right, current, goLeft, true)
	} else if !goLeft && current.rightValue != "" {
		current.rightValue = addDay18Values(sourceValue, current.rightValue)
		return true
	} else if goLeft && current.left != former && current.left != nil {
		return findAndAddToFirstInDirection(sourceValue, current.left, current, goLeft, true)
	} else if goLeft && current.leftValue != "" {
		current.leftValue = addDay18Values(sourceValue, current.leftValue)
		return true
	} else if current.parent != nil {
		return findAndAddToFirstInDirection(sourceValue, current.parent, current, goLeft, descending)
	} else {
		return false
	}
}

func explodeDay18Recursive(current *day18Node, depth int) bool {
	if depth >= 5 && current.leftValue != "" && current.rightValue != "" {
		// Explode
		parent := current.parent

		findAndAddToFirstInDirection(current.leftValue, parent, current, true, false)
		findAndAddToFirstInDirection(current.rightValue, parent, current, false, false)

		if current == parent.left {
			parent.left = nil
			parent.leftValue = "0"
		} else {
			parent.right = nil
			parent.rightValue = "0"
		}

		return true
	}

	if current.left != nil {
		isLeftReduced := explodeDay18Recursive(current.left, depth+1)

		if isLeftReduced {
			return true
		}
	} else if isDay18ValueTooBig(current.leftValue) {
		split := splitDay18Value(current.leftValue)
		current.leftValue = ""
		current.left = split
		return true
	}

	if current.right != nil {
		isRightReduced := explodeDay18Recursive(current.right, depth+1)
		if isRightReduced {
			return true
		}
	} else if isDay18ValueTooBig(current.rightValue) {
		split := splitDay18Value(current.rightValue)
		current.rightValue = ""
		current.right = split
		return true
	}

	return false
}

func reduceDay18(input *day18Node) *day18Node {
	for {
		hasExploded := explodeDay18Recursive(input, 1)
		if !hasExploded {
			break
		}
	}

	return input
}
