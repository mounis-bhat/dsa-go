package stack

import "errors"

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Display() []int {
	return s.elements
}

var ErrStackEmpty = errors.New("Stack is empty")

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrStackEmpty
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}
