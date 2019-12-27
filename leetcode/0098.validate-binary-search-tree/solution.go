package problem0098

import (
	tree "github.com/Anson5555555Lee/golang_algorithms/tree-graph"
	"github.com/golang-collections/collections/stack"
)

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
