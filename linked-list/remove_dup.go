package list

// Using hash table to track duplicates.

func removeDup(l *LinkedList) {
	seen := make(map[int]struct{})
	node := l.Head
	prev := nil

	for node.Next != nil {
		if _, ok := seen[(node.Value).(int)]; ok {
			l.RemoveByValue(node.Value)
		} else {
			seen[(node.Value).(int)] = struct{}{}
			node = node.Next
		}
	}
}
