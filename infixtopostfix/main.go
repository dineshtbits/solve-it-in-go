package main

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(ele string) {
	s.items = append(s.items, ele)
}

func (s *Stack) Pop() string {
	l := len(s.items)
	ele := s.items[l-1]
	s.items = s.items[:l-1]
	return ele
}

func (s *Stack) Top() string {
	l := len(s.items)
	return s.items[l-1]
}

func (s *Stack) Size() int {
	return len(s.items)
}

func convertToPostFix(input string) string {
	result := ""
	s := &Stack{}
	for _, l := range input {
		letter := string(l)
		if letter == "+" || letter == "-" || letter == "(" || letter == "*" {

			for s.Size() > 0 && s.Top() != "(" {
				pop := s.Pop()
				result = result + pop
			}
			if s.Size() > 0 && s.Top() == "(" {
				s.Pop()
			}

			s.Push(letter)
		} else if letter == ")" {
			if s.Size() > 0 && s.Top() != "(" {
				pop := s.Pop()
				result = result + pop
			}
			if s.Size() > 0 && s.Top() == "(" {
				s.Pop()
			}
		} else {
			result = result + letter
		}
	}
	for s.Size() > 0 {
		result = result + s.Pop()
	}
	return result
}

func main() {
	input := "(A+B)*C-D"
	result := convertToPostFix(input)
	fmt.Printf("Postfix ecpression of %v is %v\n", input, result)
}
