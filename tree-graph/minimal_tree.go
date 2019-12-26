package tg

// Given a sorted (increasing order) array with unique integer elements
// Write an algorithm to create a binary search tree with minimal height
// Hints: #19, #73, #116

func SortedArrayToBST(input []int) *BinaryTreeNode {
	// Base case
	if len(input) == 0 {
		return nil
	}

	mid := len(input) / 2

	return &BinaryTreeNode{
		Value: input[mid],
		Left:  SortedArrayToBST(input[0:mid]),
		Right: SortedArrayToBST(input[mid+1:]),
	}
}
