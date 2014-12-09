package LinkedList

type List struct {
	head *node
	end  *node
}

func NewList() *List {
	var l List = List{}

	return &l
}

func (l *List) Append(n int) {
	if l.head == nil {
		l.head = NewNode()
		l.end = NewNode()
		l.head.Data = n
	} else {
		temp := l.head

		for temp.Next != nil {
			temp = temp.Next
		}

		node := NewNode()
		node.Data = n
		node.Prev = temp

		temp.Next = node
		l.end = node
	}
}

func (l *List) Delete(n int) error {
	if l.head != nil {
		temp := l.head

		for temp != nil {
			if temp.Next != nil && temp.Next.Data == n {
				temp.Next = temp.Next.Next

				if temp.Next != nil {
					temp.Next.Prev = temp
				}

				break
			}

			temp = temp.Next
		}
	}

	return nil
}

func (l *List) ListRev() []int {
	arr := make([]int, 1, 10)

	if l.head != nil {
		temp := l.end
		arr[0] = temp.Data

		for temp.Prev != nil {
			temp = temp.Prev
			arr = append(arr, temp.Data)
		}
	}

	return arr
}

func (l *List) List() []int {
	arr := make([]int, 1, 10)

	if l.head != nil {
		temp := l.head
		arr[0] = temp.Data

		for temp.Next != nil {
			temp = temp.Next
			arr = append(arr, temp.Data)
		}
	}

	return arr
}
