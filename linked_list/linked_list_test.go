package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	type Student struct {
		Name string
		Age  int
	}

	var a Student

	fmt.Println(a, a.Name, a.Age, a.Name == "")

	list := NewLinkedList()
	fmt.Println(list.Head)

	list.Prepend(1)

	fmt.Println(list.Head.Value)
	fmt.Println(list.Head)
}
