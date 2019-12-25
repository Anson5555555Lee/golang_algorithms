package snq

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	if s.IsEmpty() != true {
		t.Fatal("wrong")
	}

	s.Push(1)
	if s.IsEmpty() == true {
		t.Fatal("wrong")
	}
	if s.stack[0] != 1 {
		t.Fatal("wrong")
	}

	s.Push(2)
	s.Push(3)
	if s.stack[0] != 3 || s.stack[1] != 2 || s.stack[2] != 1 {
		t.Fatal("wrong")
	}

	for i := 3; i > 0; i-- {
		if s.Peek() != i {
			t.Fatal("wrong")
		}
		actual := s.Pop()
		if actual != i {
			t.Fatal("wrong")
		}
	}

	if s.IsEmpty() != true {
		t.Fatal("wrong")
	}

}
