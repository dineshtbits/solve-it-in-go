package main

import (
	"fmt"
	"strconv"
)

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

func eval(op1, op2, e string) int {
	if e == "+" {
		l, _ := strconv.Atoi(op1)
		r, _ := strconv.Atoi(op2)
		return l + r
	} else if e == "-" {
		l, _ := strconv.Atoi(op1)
		r, _ := strconv.Atoi(op2)
		return l - r
	}
	return 0
}

func evalPrefix(prefix string) int {
	operatorStack := &Stack{}
	operandStack := &Stack{}
	var finalResult int

	for _, s := range prefix {
		str := string(s)
		if str == ")" {
			var result int
			for operandStack.Size() > 0 && operandStack.Top() != "(" {
				fmt.Println("here")
				op1 := operandStack.Pop()
				e := operatorStack.Pop()
				op2 := operandStack.Pop()
				result = eval(op1, op2, e)
				fmt.Printf("result: %v\n", result)
			}
			if operandStack.Size() > 0 && operandStack.Top() == "(" {
				operandStack.Pop()
			}
			operandStack.Push(strconv.Itoa(result))
		} else if str == "+" || str == "-" {
			operatorStack.Push(str)
		} else {
			operandStack.Push(str)
		}
	}

	fmt.Println(operandStack.items)
	for operatorStack.Size() > 0 {
		op1 := operandStack.Pop()
		e := operatorStack.Pop()
		op2 := operandStack.Pop()
		finalResult = eval(op2, op1, e)
	}

	return finalResult
}

func main() {
	prefix := "(1+3)-(3+8)"
	fmt.Printf("Eval of %v is %v\n", prefix, evalPrefix(prefix))
}
