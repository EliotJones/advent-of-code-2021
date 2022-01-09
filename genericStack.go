package main

type genericStack []interface{}

func (s *genericStack) Push(val interface{}) *genericStack {
	*s = append(*s, val)
	return s
}

func (s *genericStack) Pop() (*genericStack, interface{}) {
	length := len(*s)
	if length == 0 {
		return s, nil
	}

	res := (*s)[length-1]
	*s = (*s)[:length-1]

	return s, res
}

func (s genericStack) Peek() interface{} {
	length := len(s)
	if length == 0 {
		return nil
	}

	return s[length-1]
}
