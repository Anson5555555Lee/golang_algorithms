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
	if this == nil {
		return nil
	}

	result := []int{}
	nodes := stack.New()

	for {
		for this != nil {
			nodes.Push(this)
			this = this.Left
		}

		if nodes.Len() == 0 {
			break
		}

		this = (nodes.Pop()).(*BinaryTreeNode)
		result = append(result, this.Value)
		this = this.Right
	}

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

func (this *BinaryTreeNode) PreOrderIterative() []int {
	if this == nil {
		return nil
	}

	result := []int{}
	nodes := stack.New()

	// handle node value
	nodes.Push(this)

	for nodes.Len() > 0 {
		node := (nodes.Pop()).(*BinaryTreeNode)
		result = append(result, node.Value)

		// Right child is pushed first so that during pop, left is processed first.
		if node.Right != nil {
			nodes.Push(node.Right)
		}

		if node.Left != nil {
			nodes.Push(node.Left)
		}
	}

	return result
}

func (this *BinaryTreeNode) PostOrder() []int {
	if this == nil {
		return nil
	}

	result := []int{}

	if this.Left != nil {
		result = append(result, this.Left.PostOrder()...)
	}

	if this.Right != nil {
		result = append(result, this.Right.PostOrder()...)
	}

	result = append(result, this.Value)

	return result
}

func (this *BinaryTreeNode) PostorderIterativeTwoStacks() []int {
	if this == nil {
		return nil
	}

	result := []int{}
	nodes1 := stack.New()
	nodes2 := stack.New()

	nodes1.Push(this)

	for nodes1.Len() != 0 {
		node := (nodes1.Pop()).(*BinaryTreeNode)
		nodes2.Push(node)

		if node.Left != nil {
			nodes1.Push(node.Left)
		}

		if node.Right != nil {
			nodes1.Push(node.Right)
		}
	}

	for nodes2.Len() != 0 {
		result = append(result, (nodes2.Pop()).(*BinaryTreeNode).Value)
	}

	return result
}

// However, note that the above solution copies the entire tree into a stack before it starts printing.
// So the space requirements of this solutions is O(n) where as other traversals had space requirements of O(h)
// h being the tree-height. An O(h) solution for post-order is also given below.
// (And it requires only a single stack!)

func (this *BinaryTreeNode) PostorderIterativeSingleStack() []int {

	if this == nil {
		return nil
	}

	result := []int{}

	nodes := stack.New()
	var prev *BinaryTreeNode

	nodes.Push(this)

	for nodes.Len() != 0 {
		current := (nodes.Peek()).(*BinaryTreeNode)
		// Just think about going down the tree in this "if"
		// block, either left side or the right side
		//  (but not both sides)
		// If a node has both left and right children, this "if" block
		// will put only the left in the stack
		if prev == nil || prev.Left == current || prev.Right == current {
			if current.Left != nil {
				nodes.Push(current.Left)
			} else if current.Right != nil {
				nodes.Push(current.Right)
			} else {
				result = append(result, (nodes.Pop()).(*BinaryTreeNode).Value)
			}
		} else if current.Left == prev { //coming up from the left side
			if current.Right != nil {
				nodes.Push(current.Right)
			} else {
				result = append(result, (nodes.Pop()).(*BinaryTreeNode).Value)
			}
		} else if current.Right == prev { // coming up from the right side
			result = append(result, (nodes.Pop()).(*BinaryTreeNode).Value)

		}

		prev = current
	}

	return result
}
