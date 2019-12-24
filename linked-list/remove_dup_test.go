package list

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDup(t *testing.T) {
	linkedList := NewLinkedList()
	vals := []int{1, 1, 1, 2, 3, 3, 4, 5, 5, 6}
	linkedList.FromArray(vals)

	removeDup(linkedList.Head)
	actual := linkedList.Slice()
	expected := []int{1, 2, 3, 4, 5, 6}

	for _, v := range actual {
		fmt.Println(v)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("wrong")
	}
}

func TestRemoveDupRunnerSolution(t *testing.T) {
	linkedList := NewLinkedList()
	vals := []int{1, 1, 1, 2, 3, 3, 4, 5, 5, 6}
	linkedList.FromArray(vals)

	removeDupRunnerSolution(linkedList.Head)
	actual := linkedList.Slice()
	expected := []int{1, 2, 3, 4, 5, 6}

	for _, v := range actual {
		fmt.Println(v)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("wrong")
	}
}
