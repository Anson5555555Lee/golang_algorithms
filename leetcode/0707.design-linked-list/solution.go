package problem0707

type MyLinkedList struct {
	Head, Tail *Node
	Size       int
}

type Node struct {
	Value int
	Next  *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{nil, nil, 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.Head == nil || index >= this.Size {
		return -1
	}

	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	return curr.Value
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{val, nil}

	if this.Head == nil {
		this.Head = newNode
		this.Tail = newNode
	} else {
		newNode.Next = this.Head
		this.Head = newNode
	}

	this.Size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{val, nil}

	if this.Tail == nil {
		this.Head = newNode
		this.Tail = newNode
	} else {
		this.Tail.Next = newNode
		this.Tail = newNode
	}

	this.Size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.Size {
		this.AddAtTail(val)
	} else if index > this.Size {
		return
	} else {
		curr := this.Head
		count := 0
		var prev *Node
		for curr != nil {
			if count == index {
				newNode := &Node{val, nil}
				newNode.Next = curr
				prev.Next = newNode
				break
			}

			prev = curr
			curr = curr.Next
			count++
		}

		this.Size++
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.Size {
		return
	}

	var prev *Node
	curr := this.Head
	count := 0

	for curr != nil {
		if count == index {
			// delete at head
			if prev == nil {
				if this.Head == this.Tail {
					this.Head = nil
					this.Tail = nil
					this.Size = 0
					return
				}

				this.Head = curr.Next
				this.Size--
				return
			} else if index == this.Size-1 {
				this.Tail = prev
				this.Size--
				return
			} else {
				prev.Next = curr.Next
				this.Size--
			}

			break
		}

		prev = curr
		curr = curr.Next
		count++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
