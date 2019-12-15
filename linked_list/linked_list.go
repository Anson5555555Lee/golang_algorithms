package linked_list

type Node struct {
	Value int
	Next  *Node
}

func NewNode(v int) *Node {
	return &Node{Value: v}
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Prepend(val int) {
	newNode := NewNode(val)

	newNode.Next = l.Head

	l.Head = newNode
}
