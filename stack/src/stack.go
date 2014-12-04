package Stack

import "fmt"

const Max int = 10

type Stack struct {
	Arr [Max]int
	Top int
}

func (s *Stack) Push(n int) {
	if s.Top < Max {
		s.Arr[s.Top] = n
		s.Top++
	} else {
		fmt.Printf("Stack is full!\n")
	}
}

func (s *Stack) Pop() int {
	var num int

	num = s.Arr[s.Top-1]
	num = 2

	return num
}
