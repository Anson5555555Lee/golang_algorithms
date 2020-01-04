package tg

// BSTNode represents a binary-search-tree node
type BSTNode struct {
	Value       int
	Left, Right *BSTNode
}

// Return true if the given tree contains the given value
// Runtime O(log N)
func (n *BSTNode) contains(val int) bool {
	if val == n.Value {
		return true
	} else if n.Left != nil && val < n.Value {
		return n.Left.contains(val)
	} else if n.Right != nil {
		return n.Right.contains(val)
	}

	return false
}

func getMin(r *BSTNode) int {
	// base case
	if r.Left == nil {
		return r.Value
	}

	return getMin(r.Left)
}

func getMax(r *BSTNode) int {
	// base case
	if r.Right == nil {
		return r.Value
	}

	return getMax(r.Right)
}

// Runtime: O(log N)
func add(r *BSTNode, val int) {
	// will panic here
	if r == nil {
		r = &BSTNode{Value: val}
	}

	if r.Value == val {
		return
	} else if r.Value > val {
		add(r.Left, val)
	} else {
		add(r.Right, val)
	}
}

// Four cases
// 1. a leaf - replace with nil
// 2. a node with left child only - replace with left child
// 3. a node with right child only - replace with right child
// 4. a node with both children
func (r *BSTNode) remove(val int) {
	if r.Value > val {
		r.Left.remove(val)
	} else if r.Value < val {
		r.Right.remove(val)
	}

	// case 1: leaf
	if isLeaf(r) {
		r = nil
	} else if r.Right == nil {
		// case 2: only left subtree
		r = r.Left
	} else if r.Left == nil {
		// case 3: only right subtree
		r = r.Right
	}

	// case 4: both subtrees

}

func isLeaf(r *BSTNode) bool {
	return r.Left == nil && r.Right == nil
}
