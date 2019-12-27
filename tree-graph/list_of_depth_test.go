package tg

import (
	"fmt"
	"testing"
)

func TestSolutionDFS(t *testing.T) {
	root := &BinaryTreeNode{Value: 1}
	left := &BinaryTreeNode{Value: 3}
	right := &BinaryTreeNode{Value: 2}
	grandLeft := &BinaryTreeNode{Value: 5}
	grandRight := &BinaryTreeNode{Value: 6}
	grandGrandLeft := &BinaryTreeNode{Value: 7}

	root.Left = left
	root.Right = right
	left.Left = grandLeft
	left.Right = grandRight
	grandLeft.Left = grandGrandLeft

	actual := SolutionDFS(root)
	for _, l := range actual {
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
		fmt.Println()
	}
}
