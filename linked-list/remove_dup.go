package list

// Using hash table to track duplicates.
// Solve the problem with one pass.
func removeDup(n *Node) {
	seen := make(map[int]struct{})
	prev := &Node{}

	for n != nil {
		if _, ok := seen[(n.Value).(int)]; ok {
			prev.Next = n.Next
			// n.Next = nil
		} else {
			seen[(n.Value).(int)] = struct{}{}
			prev = n
		}

		n = n.Next
	}
}

// No buffer allowed
// If we don't have a buffer, we can iterates through the linked list
// and runner which checks all subsequent nodes for duplicates
// this code runs in O(1) space, but O(N^2) time.
func removeDupRunnerSolution(n *Node) {
	curr := n

	for curr != nil {
		// remove all future nodes that have the same value
		runner := curr
		for runner.Next != nil {
			if runner.Next.Value == curr.Value {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}

		curr = curr.Next
	}
}
