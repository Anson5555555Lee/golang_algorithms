package tg

import (
	"reflect"
	"testing"
)

func TestBinaryTreeNode(t *testing.T) {
	node := BinaryTreeNode{}
	if node.Value != 0 {
		t.Fatal("wrong")
	}

	if node.Left != nil {
		t.Fatal("wrong")
	}

	if node.Right != nil {
		t.Fatal("wrong")
	}

	leftNode := &BinaryTreeNode{Value: 1}
	rightNode := &BinaryTreeNode{Value: 3}
	rootNode := &BinaryTreeNode{Value: 2}

	rootNode.Left = leftNode
	rootNode.Right = rightNode

	if rootNode.Left.Value != 1 {
		t.Fatal("wrong")
	}

	if rootNode.Right.Value != 3 {
		t.Fatal("wrong")
	}
}

func TestBinaryTreeNodeHeight(t *testing.T) {
	root := &BinaryTreeNode{Value: 1}
	left := &BinaryTreeNode{Value: 3}
	right := &BinaryTreeNode{Value: 2}
	grandLeft := &BinaryTreeNode{Value: 5}
	grandRight := &BinaryTreeNode{Value: 6}
	grandGrandLeft := &BinaryTreeNode{7, nil, nil}

	if root.GetHeight() != 0 || root.BalanceFactor() != 0 {
		t.Fatal("wrong")
	}

	root.Left = left
	root.Right = right
	if root.GetHeight() != 1 || left.GetHeight() != 0 || right.GetHeight() != 0 || root.BalanceFactor() != 0 {
		t.Fatal("wrong")
	}

	left.Left = grandLeft
	left.Right = grandRight
	if root.GetHeight() != 2 ||
		left.GetHeight() != 1 ||
		right.GetHeight() != 0 ||
		grandLeft.GetHeight() != 0 ||
		grandRight.GetHeight() != 0 ||
		root.BalanceFactor() != 1 {
		t.Fatal("wrong")
	}

	grandLeft.Left = grandGrandLeft

	// In-order traversal
	actual := root.InOrder()
	expected := []int{7, 5, 3, 6, 1, 2}
	if reflect.DeepEqual(actual, expected) != true {
		t.Fatal("wrong")
	}

	// Pre-order traversal
	actual = root.PreOrder()
	expected = []int{1, 3, 5, 7, 6, 2}
	if reflect.DeepEqual(actual, expected) != true {
		t.Fatal("wrong")
	}
}
