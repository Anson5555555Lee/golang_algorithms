package list

type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode(v interface{}) *Node {
	return &Node{Value: v}
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Prepend(val interface{}) *LinkedList {
	newNode := NewNode(val)
	newNode.Next = l.Head

	l.Head = newNode

	return l
}

func (l *LinkedList) Append(val interface{}) *LinkedList {
	newNode := NewNode(val)

	if l.Head == nil {
		l.Head = newNode

		return l
	}

	currNode := l.Head
	// for currNode != nil {
	// 	if currNode.Next != nil {
	// 		currNode = currNode.Next
	// 		continue
	// 	}

	// 	currNode.Next = newNode
	// 	break
	// }
	for currNode.Next != nil {
		currNode = currNode.Next
	}

	currNode.Next = newNode

	return l
}

func (l *LinkedList) RemoveByValue(val interface{}) bool {
	if l.Head == nil {
		return false
	}

	if l.Head.Value == val {
		l.Head = l.Head.Next
		return true
	}

	currNode := l.Head
	for currNode.Next != nil {
		if currNode.Next.Value == val {
			currNode.Next = currNode.Next.Next

			return true
		}

		currNode = currNode.Next
	}

	return false
}

func (l *LinkedList) FindByValue(val interface{}) *Node {
	if l.Head == nil {
		return nil
	}

	currNode := l.Head
	for currNode != nil {
		if currNode.Value == val {
			return currNode
		}

		currNode = currNode.Next
	}

	return nil
}

func (l *LinkedList) Reverse() {
	var prev, next *Node
	curr := l.Head

	for curr != nil {
		// Store next node.
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	l.Head = prev
}
