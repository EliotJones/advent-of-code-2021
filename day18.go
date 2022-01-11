package main

import "strings"

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
