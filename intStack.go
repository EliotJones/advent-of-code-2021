package main

type intStack []int

func (s *intStack) Push(val int) *intStack {
	*s = append(*s, val)
	return s
}

func (s *intStack) Pop() (*intStack, int) {
	length := len(*s)
	if length == 0 {
		return s, -1
	}

	res := (*s)[length-1]
	*s = (*s)[:length-1]

	return s, res
}

func (s intStack) Peek() int {
	length := len(s)
	if length == 0 {
		return -1
	}

	return s[length-1]
}
