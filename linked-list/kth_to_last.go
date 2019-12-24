package list

// Implement an algorithm to find the kth to last element of a singly linked list.
// Hints: #8, #25, #41, #67, #126

// For this solution, we have defined k that passing in k = 1 would return the last element
// k = 2 would return the second to the last element, and so on.

// Solution #1: if linked list size is known
// If the size of the linked list is known, the kth to last element is the (length - k)th element.
// We can just iterate through the linked list to find this element.

func kthToLast(head *Node, k int) *Node {
	curr := head
	l := 0
	for curr != nil {
		l++
		curr = curr.Next
	}

	curr = head
	for i := 0; i < l-k; i++ {
		curr = curr.Next
	}

	return curr
}

// Solution #2: recursive
func kthToLastRecursive(head *Node, k int) (*Node, int) {
	if head == nil {
		return head, 0
	}

	node, index := kthToLastRecursive(head.Next, k)
	index++

	if k == index {
		return head, index
	}

	return node, index
}

// Solution #3: iterative
// We can use two pointers, p1 and p2.
// We place them k nodes apart by putting p2 at the beginning and moving p1
// k nodes into the list. Then we move them at the same pace, p1 will hit the end
// of the linked list after Length - k steps.
// At that point, p2 will be length - k nodes into the list, or k nodes
// from the end.
// This method takes O(n) time and O(1) space.
func kthToLastIterative(head *Node, k int) *Node {
	p1, p2 := head, head

	for i := 0; i < k; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.Next
	}

	// Move them at the same space.
	// When p1 hits the end, p2 will be at the right element.
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}
