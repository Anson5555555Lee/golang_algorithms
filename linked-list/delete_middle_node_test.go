package list

import (
	"reflect"
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
	l := NewLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	node := l.Head
	for i := 0; i < 2; i++ {
		node = node.Next
	}

	deleteMiddle(node)
	actual := l.Slice()
	expected := []int{1, 2, 4, 5}
	if reflect.DeepEqual(actual, expected) != true {
		t.Fatal("wrong")
	}

}
