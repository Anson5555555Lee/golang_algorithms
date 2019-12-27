package tg

import "github.com/golang-collections/collections/stack"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (this *BinaryTreeNode) GetLeftHeight() int {
	if this.Left == nil {
		return 0
	}

	return this.Left.GetHeight() + 1
}

func (this *BinaryTreeNode) GetRightHeight() int {
	if this.Right == nil {
		return 0
	}

	return this.Right.GetHeight() + 1
}

func (this *BinaryTreeNode) GetHeight() int {
	if this.GetLeftHeight() > this.GetRightHeight() {
		return this.GetLeftHeight()
	}

	return this.GetRightHeight()
}

func (this *BinaryTreeNode) BalanceFactor() int {
	return this.GetLeftHeight() - this.GetRightHeight()
}

func (this *BinaryTreeNode) InOrder() []int {
	result := []int{}

	// Add left nodes
	if this.Left != nil {
		result = append(result, this.Left.InOrder()...)
	}

	// Add root.
	result = append(result, this.Value)

	// Add right nodes.
	if this.Right != nil {
		result = append(result, this.Right.InOrder()...)
	}

	return result
}

func (this *BinaryTreeNode) InOrderIterative() []int {
	result := []int{}

	nodes := stack.New()

	return result
}

func (this *BinaryTreeNode) PreOrder() []int {
	result := []int{}

	// Add root
	result = append(result, this.Value)

	// Add left nodes
	if this.Left != nil {
		result = append(result, this.Left.PreOrder()...)
	}

	// Add right nodes
	if this.Right != nil {
		result = append(result, this.Right.PreOrder()...)
	}

	return result
}
