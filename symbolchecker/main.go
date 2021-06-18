package main

import "fmt"

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(elem interface{}) {
	s.elements = append(s.elements, elem)
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Pop() interface{} {
	l := len(s.elements)
	elem := s.elements[l-1]
	s.elements = s.elements[:l-1]
	return elem
}

func (s *Stack) Top() interface{} {
	l := len(s.elements)
	return s.elements[l-1]
}

func BalancedSymbols(s string) bool {
	myStack := &Stack{}
	for _, c := range s {

		if string(c) == "{" || string(c) == "(" || string(c) == "[" {
			// fmt.Printf("Pushing %v\n", string(c))
			myStack.Push(string(c))
		} else if string(c) == "}" || string(c) == ")" || string(c) == "]" {
			if (myStack.Top().(string) == "{" && string(c) == "}") || (myStack.Top().(string) == "(" && string(c) == ")") || (myStack.Top().(string) == "[" && string(c) == "]") {
				myStack.Pop()
			}
		}
	}
	return myStack.Size() <= 0
}

func main() {
	s := "(A+B)+(C+D)"
	fmt.Printf("%v is balanced? %v\n", s, BalancedSymbols(s))
	s = "(A+B)+(C+D"
	fmt.Printf("%v is balanced? %v\n", s, BalancedSymbols(s))
}
