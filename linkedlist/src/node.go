package LinkedList

type node struct {
	Data int
	Link *node
}

func NewNode() *node {
	var n node = node{}
	n.Data = -1

	return &n
}
