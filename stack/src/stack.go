package Stack

const max int = 10

type Stack struct {
	arr [max]int
	top int
}

func NewStack() *Stack {
	var s Stack = Stack{}
	s.top = -1

	return &s
}

func (s *Stack) Size() int {
	return s.top
}

func (s *Stack) Push(n int) bool {
	status := false

	if s.top < max {
		s.top++
		s.arr[s.top] = n

		status = true
	}

	return status
}

func (s *Stack) Pop() int {
	num := -1

	if s.top >= 0 {
		num = s.arr[s.top]
		s.top--
	}

	return num
}
