package problem0098

import (
	tree "github.com/Anson5555555Lee/golang_algorithms/tree-graph"
	"github.com/golang-collections/collections/stack"
)

// for a BST
// left <= node < right
func solution(root *tree.BinaryTreeNode) bool {
	if root == nil {
		return true
	}

	nodes := stack.New()
	var prev *tree.BinaryTreeNode

	for {
		for root != nil {
			nodes.Push(root)
			root = root.Left
		}

		if nodes.Len() == 0 {
			break
		}

		root = (nodes.Pop()).(*tree.BinaryTreeNode)
		if prev != nil && prev.Value >= root.Value {
			return false
		}
		prev = root
		root = root.Right
	}

	return true
}

func solutionMinMax(root *tree.BinaryTreeNode) bool {
	// min and max values for int type
	Min, Max := -1<<63, 1<<63-1

	return recur(Min, Max, root)
}

func recur(min, max int, node *tree.BinaryTreeNode) bool {
	if node == nil {
		return true
	}

	if node.Value <= min || node.Value >= max {
		return false
	}

	if recur(min, node.Value, node.Left) == false || recur(node.Value, max, node.Right) == false {
		return false
	}

	return true
}
