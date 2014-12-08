package Queue

const max int = 10

type Queue struct {
	arr   []int
	front int
	end   int
}

func NewQueue() *Queue {
	var q Queue = Queue{}
	q.arr = make([]int, 1, max)
	q.front = -1
	q.end = -1

	return &q
}

func (q *Queue) Front() int {
	return (q.front)
}

func (q *Queue) End() int {
	return (q.end)
}

func (q *Queue) Len() int {
	return len(q.arr)
}

func (q *Queue) Cap() int {
	return cap(q.arr)
}

func (q *Queue) Add(n int) {
	if q.front == -1 && q.end == -1 {
		q.front++
	}

	q.end++

	q.arr = append(q.arr, n)
}

func (q *Queue) Del() int {
	num := -1
	temp := -1

	num = q.front

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
