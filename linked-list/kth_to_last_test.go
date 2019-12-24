package list

import "testing"

func TestKthToLast(t *testing.T) {
	l := NewLinkedList()

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	k := 1
	actual := kthToLast(l.Head, k).Value
	expected := 5
	if actual != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual)
	}

	k = 3
	actual = kthToLast(l.Head, k).Value
	expected = 3
	if actual != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual)
	}
}

func TestKthToLastRecursive(t *testing.T) {
	l := NewLinkedList()

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	k := 1
	actual, _ := kthToLastRecursive(l.Head, k)
	expected := 5
	if actual.Value != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual.Value)
	}

	k = 3
	actual, _ = kthToLastRecursive(l.Head, k)
	expected = 3
	if actual.Value != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual.Value)
	}

	k = 5
	actual, _ = kthToLastRecursive(l.Head, k)
	expected = 1
	if actual.Value != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual.Value)
	}
}

func TestKthToLastIterative(t *testing.T) {
	l := NewLinkedList()

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	k := 1
	actual := kthToLastIterative(l.Head, k).Value
	expected := 5
	if actual != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual)
	}

	k = 3
	actual = kthToLastIterative(l.Head, k).Value
	expected = 3
	if actual != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual)
	}

	k = 5
	actual = kthToLastIterative(l.Head, k).Value
	expected = 1
	if actual != expected {
		t.Errorf("%dth element expected to be %v, got %v", k, expected, actual)
	}
}
