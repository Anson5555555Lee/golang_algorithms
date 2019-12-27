package tg

import "container/list"

// Given a binary tree, design an algorithm which creates a linked list of all the nodes at each depth
// e.g., if you have a tree with depth D, you'll have D linked lists
// #107, #123, #135
func SolutionDFS(root *BinaryTreeNode) []*list.List {
	result := []*list.List{}

	var dfs func(*BinaryTreeNode, int)

	dfs = func(root *BinaryTreeNode, level int) {
		if root == nil {
			return
		}

		// found new level
		if level == len(result) {
			result = append(result, list.New())
		}
		result[level].PushBack(root.Value)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)

	}

	dfs(root, 0)

	return result
}
