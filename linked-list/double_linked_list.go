package list

type node struct {
	value int
	next  *node
	prev  *node
}

type DoubleLinkedList struct {
	head  *node
	tail  *node
	count int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{nil, nil, 0}
}

func GetDoubleLinkedListFromValues(vals []int) *DoubleLinkedList {
	l := NewDoubleLinkedList()

	if len(vals) == 0 {
		return l
	}

	for _, val := range vals {
		l.insert(val)
	}

	return l
}

// append value at the end
func (dl *DoubleLinkedList) insert(val int) {
	newNode := &node{value: val}
	dl.insertNode(newNode)
}

func (dl *DoubleLinkedList) insertNode(newNode *node) {
	dl.count++
	if dl.tail == nil {
		dl.head, dl.tail = newNode, newNode
	} else {
		dl.tail.next = newNode
		newNode.prev = dl.tail
		dl.tail = newNode
	}
}

func (dl *DoubleLinkedList) getNode(index int) *node {
	node := dl.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

// Mostly used for testing.

func (dl *DoubleLinkedList) len() int {
	return dl.count
}
