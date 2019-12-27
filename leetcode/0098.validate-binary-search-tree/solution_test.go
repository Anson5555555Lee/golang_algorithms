package problem0098

import (
	"testing"

	tree "github.com/Anson5555555Lee/golang_algorithms/tree-graph"
)

func TestSolution(t *testing.T) {
	root := &tree.BinaryTreeNode{Value: 1}
	left := &tree.BinaryTreeNode{Value: 3}
	right := &tree.BinaryTreeNode{Value: 2}
	grandLeft := &tree.BinaryTreeNode{Value: 5}
	grandRight := &tree.BinaryTreeNode{Value: 6}
	grandGrandLeft := &tree.BinaryTreeNode{Value: 7}

	root.Left = left
	root.Right = right
	left.Left = grandLeft
	left.Right = grandRight
	grandLeft.Left = grandGrandLeft

	actual := solution(root)
	expected := false
	if actual != expected {
		t.Fatal("wrong")
	}

	root2 := &tree.BinaryTreeNode{Value: 2}
	left2 := &tree.BinaryTreeNode{Value: 1}
	right2 := &tree.BinaryTreeNode{Value: 3}
	root2.Left = left2
	root2.Right = right2

	actual = solution(root2)
	expected = true
	if actual != expected {
		t.Fatal("wrong")
	}
}
