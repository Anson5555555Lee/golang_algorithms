package snq

import "testing"

func TestQueueUsingStacks(t *testing.T) {
	q := NewQueueUsingStacks()

	q.Add(1)
	q.Add(2)
	q.Add(3)

	actual, _ := q.Remove()
	if actual != 1 {
		t.Fatal("wrong")
	}

	q.Add(4)
	q.Add(5)
	actual, _ = q.Remove()
	if actual != 2 {
		t.Fatal("wrong")
	}
}
