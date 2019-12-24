package list

// Implement an algorithm to delete a node in the middle of a singly linked list
// i.e., any node but the first and last node, not necessarily the exact middle
// Hints: #72

func deleteMiddle(n *Node) bool {
	if n.Next == nil {
		return false
	}

	// next := n.Next

	n.Value = n.Next.Value
	n.Next = n.Next.Next

	return true
}
