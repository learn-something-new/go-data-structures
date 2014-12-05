package Stack

const max int = 10

type Stack struct {
	arr []int
	top int
}

func NewStack() *Stack {
	var s Stack = Stack{}
	s.arr = make([]int, max)
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

func (s *Stack) Push(n int) {
	s.top++

	if (s.top) >= cap(s.arr) {
		newArr := make([]int, max)
		s.arr = append(s.arr, newArr...)
	}

	s.arr[s.top] = n
}

func (s *Stack) Pop() int {
	num := -1

	if s.top >= 0 {
		num = s.arr[s.top]
		s.top--
	}

	return num
}
