package ds

import "errors"

type Stack struct {
	data     []int
	size     int
	capacity int
}

func MakeStack() Stack {
	return Stack{
		data: make([]int, 10),
		size: 0,
		capacity: 10,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(value int) {
	if s.capacity == s.size {
		s.resize()
	}

	s.data[s.size] = value
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("tried to pop from an empty stack")
	}
	
	v := s.data[s.size - 1]
	s.size--;
	return v, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("tried to peek from an empty stack")
	}

	return s.data[s.size - 1], nil
}

func (s *Stack) resize() {
	data := make([]int, s.capacity * 2)
	for i, v := range s.data {
		data[i] = v
	}
	s.data = data
	s.capacity *= 2
}
