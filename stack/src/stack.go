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

func (s *Stack) Size() int {
	return s.top
}

func (s *Stack) Push(n int) bool {
	status := false

	if (s.top + 1) < cap(s.arr) {
		s.top++
		s.arr[s.top] = n

		status = true
	} else {
		newArr := make([]int, (len(s.arr)+1)*2)
		copy(newArr, s.arr)
		s.arr = newArr

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
