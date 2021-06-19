package main

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

func main() {

}
