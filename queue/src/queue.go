package Queue

const max int = 10

type Queue struct {
	arr   []int
	front int
	end   int
}

func NewQueue() *Queue {
	var q Queue = Queue{}
	q.arr = make([]int, max)
	q.front = -1
	q.end = -1

	return &q
}

func (q *Queue) Add(n int) {
	if q.front == -1 && q.end == -1 {
		q.front++
		q.end++

		q.arr[q.end] = n
	} else {
		q.end++

		if q.end < max {
			q.arr[q.end] = n
		} else {
			newArr := make([]int, (cap(q.arr)+1)*2)
			copy(newArr, q.arr)
			q.arr = newArr
		}
	}
}

func (q *Queue) Del() int {
	num := -1
	temp := -1

	if q.front >= -1 {
		num = q.arr[q.front]

		for x := 0; x <= q.end; x++ {

			if x+1 <= q.end {
				temp = q.arr[x+1]
				q.arr[x] = temp
			} else {
				q.end--

				if q.end == -1 {
					q.front = q.end
				} else {
					q.front = 0
				}
			}
		}
	}

	return num
}
