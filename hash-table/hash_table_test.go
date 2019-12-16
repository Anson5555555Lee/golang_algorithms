package ht

import (
	"testing"
)

func TestHashTableCreateSpecificSize(t *testing.T) {
	ht := New(32)
	if ht.Capacity != 32 {
		t.Fatalf("Expected %d but got %d", 32, ht.Capacity)
	}
	if ht.Size != 0 {
		t.Fatalf("Expected %d but got %d", 0, ht.Size)
	}
}

func TestHashFunction(t *testing.T) {
	bucketSize := 32
	h := hash("a") % bucketSize
	if h != 1 {
		t.Fatalf("Expected %d but got %d", 1, h)
	}

	h = hash("b") % bucketSize
	if h != 2 {
		t.Fatalf("Expected %d but got %d", 2, h)
	}

	h = hash("abc") % bucketSize
	if h != 6 {
		t.Fatalf("Expected %d but got %d", 6, h)
	}
}
