package NextedList

type List struct {
	p *node
}

func NewList() *List {
	var l List = List{}

	return &l
}

func (l *List) Append(n int) {
	if l.p == nil {
		l.p = NewNode()
		l.p.Data = n
	} else {
		temp := l.p

		for temp.Next != nil {
			temp = temp.Next
		}

		node := NewNode()
		node.Data = n

		temp.Next = node
	}
}

func (l *List) List() []int {
	arr := make([]int, 1, 10)

	if l.p != nil {
		temp := l.p
		arr[0] = temp.Data

		for temp.Next != nil {
			temp = temp.Next
			arr = append(arr, temp.Data)
		}
	}

	return arr
}
