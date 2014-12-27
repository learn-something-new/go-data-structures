package Stack

import "errors"

const max int = 10

type Stack struct {
	arr []int
	top int
}

func NewStack() *Stack {
	var s Stack = Stack{}
	s.arr = make([]int, 1, max)
	s.top = -1

	return &s
}

func (s *Stack) Top() int {
	return s.top
}

func (s *Stack) Len() int {
	return len(s.arr)
}

func (s *Stack) Cap() int {
	return cap(s.arr)
}

func (s *Stack) Push(n int) error {
	s.top++
	s.arr = append(s.arr, n)

	return nil
}

func (s *Stack) Pop() (int, error) {
	num := -1

	if s.top >= 0 {
		num = s.arr[s.top]
		s.top--
	} else {
		return -1, errors.New("Stack is empty")
	}

	return num, nil
}
