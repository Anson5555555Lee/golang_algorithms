package snq

import "testing"

func TestMinStack(t *testing.T) {
	s := NewMinStack()

	s.Push(-2)
	s.Push(0)
	s.Push(-3)

	actual := s.GetMin()
	expected := -3
	if actual != expected {
		t.Fatal("wrong")
	}

	s.Pop()
	actual = s.GetMin()
	expected = -2
	if actual != expected {
		t.Fatal("wrong")
	}
}
