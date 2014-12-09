package LinkedList

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

		for temp.Link != nil {
			temp = temp.Link
		}

		node := NewNode()
		node.Data = n

		temp.Link = node
	}
}

func (l *List) Delete(n int) error {
	if l.p != nil {
		temp := l.p

		for temp != nil {
			if temp.Link != nil && temp.Link.Data == n {
				temp.Link = temp.Link.Link
				break
			}

			temp = temp.Link
		}
	}

	return nil
}

func (l *List) List() []int {
	arr := make([]int, 1, 10)

	if l.p != nil {
		temp := l.p
		arr[0] = temp.Data

		for temp.Link != nil {
			temp = temp.Link
			arr = append(arr, temp.Data)
		}
	}

	return arr
}
