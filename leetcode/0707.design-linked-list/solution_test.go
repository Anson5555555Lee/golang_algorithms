package problem0707

import "testing"

// ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
// [[],[1],[3],[1,2],[1],[1],[1]]

func TestMyLinkedList(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	if list.Head.Value != 1 ||
		list.Head.Next != nil ||
		list.Tail.Value != 1 ||
		list.Tail.Next != nil ||
		list.Size != 1 {
		t.Fatal("wrong")
	}
	list.AddAtTail(3)
	if list.Head.Value != 1 ||
		list.Head.Next != list.Tail ||
		list.Tail.Value != 3 ||
		list.Tail.Next != nil ||
		list.Size != 2 {
		t.Fatal("wrong")
	}
	list.AddAtIndex(1, 2)
	if list.Head.Value != 1 || list.Tail.Value != 3 || list.Head.Next.Value != 2 || list.Size != 3 {
		t.Fatal("wrong")
	}

	val := list.Get(1)
	if val != 2 {
		t.Fatal("wrong")
	}

	list.DeleteAtIndex(1)
	val = list.Get(1)
	if val == 2 {
		t.Fatal("wrong")
	}
}
