package list

import (
	"testing"
)

func TestLinkedListNodeCreateWithValue(t *testing.T) {
	newNode := NewNode(1)

	if newNode.Value != 1 {
		t.Fatal("wrong")
	}

	if newNode.Next != nil {
		t.Fatal("wrong")
	}

}

func TestLinkedListNodeLinkNodesTogether(t *testing.T) {
	node2 := NewNode(2)
	node1 := NewNode(1)

	node1.Next = node2
	if node2.Next != nil {
		t.Fatal("wrong answer")
	}

	if node1.Value != 1 {
		t.Fatal("wrong answer")
	}

	if node2.Value != 2 {
		t.Fatalf("wrong answer")
	}

	if node1.Next != node2 {
		t.Fatal("wrong")
	}

	if node1.Next.Value != 2 {
		t.Fatal("wrong")
	}
}

func TestLinkedListPrependNode(t *testing.T) {
	linkedList := NewLinkedList()

	linkedList.Prepend(2)
	if linkedList.Head.Value != 2 {
		t.Fatal("wrong")
	}

	linkedList.Prepend(1)
	if linkedList.Head.Value != 1 {
		t.Fatal("wrong")
	}

	if linkedList.Head.Next.Value != 2 {
		t.Fatalf("Expected %d but got %v", 2, linkedList.Head.Next.Value)
	}
}

func TestLinkedListAppendNode(t *testing.T) {
	linkedList := NewLinkedList()

	if linkedList.Head != nil {
		t.Fatalf("Expected %v but got %v", nil, linkedList.Head)
	}

	linkedList.Append(1)
	if linkedList.Head.Value != 1 {
		t.Fatalf("Expected %d but got %d", 1, linkedList.Head.Value)
	}

	linkedList.Append(2)
	if linkedList.Head.Next.Value != 2 {
		t.Fatalf("Expected %d but got %v", 2, linkedList.Head.Next.Value)
	}

	linkedList.Append(3)
	if linkedList.Head.Next.Next.Value != 3 {
		t.Fatalf("Expected %d but got %v", 3, linkedList.Head.Next.Next.Value)
	}
	if linkedList.Head.Next.Next.Next != nil {
		t.Fatalf("Expected %v but got %v", nil, linkedList.Head.Next.Next.Next)
	}
}

func TestLinkedListRemoveByValue(t *testing.T) {
	linkedList := NewLinkedList()

	ret := linkedList.RemoveByValue(1)

	if ret != false {
		t.Fatalf("Expected %t but got %t", false, ret)
	}

	linkedList.Append(1)
	ret = linkedList.RemoveByValue(1)
	if ret != true {
		t.Fatalf("Expected %t but got %t", true, ret)
	}
	if linkedList.Head != nil {
		t.Fatalf("Expected %v but got %v", nil, linkedList.Head)
	}

	linkedList.Append(1)
	if linkedList.Head.Value != 1 {
		t.Fatalf("Expected %d but got %d", 1, linkedList.Head.Value)
	}

	linkedList.Append(2)
	ret = linkedList.RemoveByValue(3)
	if ret != false {
		t.Fatalf("Expected %t but got %t", false, ret)
	}

	ret = linkedList.RemoveByValue(2)
	if ret != true {
		t.Fatalf("Expected %t but got %t", true, ret)
	}
}

func TestLinkedListFindByValue(t *testing.T) {
	linkedList := NewLinkedList()

	node := linkedList.FindByValue(1)
	if node != nil {
		t.Fatalf("Expected %v but got %v", nil, node)
	}

	linkedList.Append(1)
	node = linkedList.FindByValue(1)
	if node == nil || node.Value != 1 {
		t.Fatalf("Expected %d but got %d", 1, node.Value)
	}

	node = linkedList.FindByValue(10)
	if node != nil {
		t.Fatalf("Expected %v but got %v", nil, node)
	}
}

func TestLinkedListFromArray(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	linkedList := NewLinkedList()
	linkedList.FromArray(input)

	if linkedList.Head.Value != 1 {
		t.Fatalf("Expected %d but got %d", 1, linkedList.Head.Value)
	}

	if linkedList.Head.Next.Value != 2 {
		t.Fatalf("Expected %d but got %d", 1, linkedList.Head.Next.Value)
	}

	if linkedList.Head.Next.Next.Value != 3 {
		t.Fatalf("Expected %d but got %d", 1, linkedList.Head.Next.Next.Value)
	}
}

func TestLinkedListReverse(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	linkedList.Reverse()
	if linkedList.Head.Value != 3 {
		t.Fatalf("Expected %d but got %d", 3, linkedList.Head.Value)
	}
	if linkedList.Head.Next.Value != 2 {
		t.Fatalf("Expected %d but got %d", 2, linkedList.Head.Next.Value)
	}
	if linkedList.Head.Next.Next.Value != 1 {
		t.Fatalf("Expected %d but got %d", 1, linkedList.Head.Next.Next.Value)
	}
	if linkedList.Head.Next.Next.Next != nil {
		t.Fatalf("Expected %d but got %d", 2, linkedList.Head.Next.Value)
	}
}
