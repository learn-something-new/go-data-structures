package LinkedList

type node struct {
	Data int
	Next *node
	Prev *node
}

func NewNode() *node {
	var n node = node{}
	n.Data = -1

	return &n
}
