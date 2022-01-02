package main

import (
	"fmt"
	"strings"
	"unicode"
)

const d12StartType byte = 0
const d12SmallType byte = 1
const d12LargeType byte = 2
const d12EndType byte = 3

type day12GraphNode struct {
	id         int
	label      string
	cavernType byte
	links      []*day12GraphNode
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func parseNodeType(name string) byte {
	if name == "start" {
		return d12StartType
	} else if name == "end" {
		return d12EndType
	} else if IsUpper(name) {
		return d12LargeType
	} else {
		return d12SmallType
	}
}

func addOrCreateNode(name string, labelToNodeMap *map[string]*day12GraphNode, id int) {
	if _, ok := (*labelToNodeMap)[name]; ok {
		// Exists.
	} else {
		(*labelToNodeMap)[name] = &day12GraphNode{
			id:         id,
			label:      name,
			cavernType: parseNodeType(name),
			links:      make([]*day12GraphNode, 0),
		}
	}
}

func parseDay12InputFromFileToGraph(path string) *day12GraphNode {
	scanner, err := scannerForFile(path)
	if err != nil {
		panic(err)
	}

	var id int
	labelToNodeMap := make(map[string]*day12GraphNode, 0)
	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.Split(line, "-")

		if len(nodes) != 2 {
			continue
		}

		if nodes[0] == "dc" {
			fmt.Println("dc node")
		}

		addOrCreateNode(nodes[0], &labelToNodeMap, id)
		// Leaves gaps but we don't need contiguous.
		id++

		addOrCreateNode(nodes[1], &labelToNodeMap, id)
		id++

		n1, n2 := labelToNodeMap[nodes[0]], labelToNodeMap[nodes[1]]

		add := true
		for _, n := range n1.links {
			if n.id == n2.id {
				add = false
				break
			}
		}

		if add {
			n1.links = append(n1.links, n2)
			n2.links = append(n2.links, n1)
		}
	}

	return labelToNodeMap["start"]
}

func recursiveVisitNext(current *day12GraphNode, route []string, result *[][]string) {

	// We always visited this one.
	route = append(route, current.label)

	// Exit at the end.
	if current.cavernType == d12EndType {
		*result = append(*result, route)
		return
	}

	newRoute := make([]string, len(route))
	copy(newRoute, route)
	for _, n := range current.links {
		if n.cavernType == d12StartType {
			continue
		} else if n.cavernType == d12SmallType {
			var isRevisit bool
			for _, label := range route {
				if label == n.label {
					isRevisit = true
					break
				}
			}

			if isRevisit {
				continue
			}

			// Visit small node
			recursiveVisitNext(n, newRoute, result)
		} else {
			recursiveVisitNext(n, newRoute, result)
		}
	}
}

func day12() {
	startNode := parseDay12InputFromFileToGraph("inputs/day12.txt")

	result := make([][]string, 0)

	recursiveVisitNext(startNode, make([]string, 0), &result)

	fmt.Println("Result is", len(result))
}
