package tg

import (
	"fmt"
	"testing"
)

var containTests = []struct {
	input    int
	expected bool
}{
	{87, true},
	{42, true},
	{55, true},
	{-3, true},
	{17, false},
	{6, false},
	{-1, false},
	{0, false},
}

func getTree() *BSTNode {
	root := &BSTNode{55, nil, nil}
	left := &BSTNode{Value: 29}
	right := &BSTNode{Value: 87}
	root.Left = left
	root.Right = right
	left.Left = &BSTNode{Value: -3}
	left.Right = &BSTNode{Value: 42}
	right.Left = &BSTNode{Value: 60}
	right.Right = &BSTNode{Value: 91}

	return root
}

func TestContains(t *testing.T) {
	root := getTree()

	for _, tt := range containTests {
		actual := root.contains(tt.input)
		if actual != tt.expected {
			t.Fatal("wrong")
		}
	}
}

func TestGetMin(t *testing.T) {
	root := getTree()

	actual := getMin(root)
	if actual != -3 {
		t.Fatal("wrong")
	}
}

func TestGetMax(t *testing.T) {
	root := getTree()

	actual := getMax(root)
	if actual != 91 {
		t.Fatal("wrong")
	}
}

// func TestAdd(t *testing.T) {
// 	var r *BSTNode
// 	add(r, 55)
// 	fmt.Println(r.Value)
// }

func TestRemove(t *testing.T) {
	root := &BSTNode{Value: 60}
	root.Left = &BSTNode{Value: 55}
	root.Left.Left = &BSTNode{Value: 29}
	root.Left.Left.Right = &BSTNode{Value: 42}
	root.remove(55)
	fmt.Println(root.Value)
	fmt.Println(root.Left.Value)
	fmt.Println(root.Left.Left.Value)
}
