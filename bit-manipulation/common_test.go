package bits

import "testing"

func TestGetBit(t *testing.T) {
	tests := []struct {
		num      int
		position uint
		expected bool
	}{
		{1, 0, true},
		{1, 1, false},
		{2, 0, false},
		{2, 1, true},
		{3, 0, true},
		{3, 1, true},
		{10, 0, false},
		{10, 1, true},
		{10, 2, false},
		{10, 3, true},
	}

	for _, tt := range tests {
		actual := GetBit(tt.num, tt.position)
		if actual != tt.expected {
			t.Fatal("wrong")
		}
	}
}

func TestSetBit(t *testing.T) {
	tests := []struct {
		num      int
		position uint
		expected int
	}{
		{1, 0, 1},
		{1, 1, 3},
		{1, 2, 5},
		{10, 0, 11},
		{10, 1, 10},
		{10, 2, 14},
		{10, 3, 10},
		{10, 4, 26},
	}

	for _, tt := range tests {
		actual := SetBit(tt.num, tt.position)
		if actual != tt.expected {
			t.Fatal("wrong")
		}
	}
}

func TestClearBit(t *testing.T) {
	tests := []struct {
		num      int
		position uint
		expected int
	}{
		{1, 0, 0},
		{1, 1, 1},
		{1, 2, 1},
		{10, 0, 10},
		{10, 1, 8},
		{10, 2, 10},
		{10, 3, 2},
	}

	for _, tt := range tests {
		actual := ClearBit(tt.num, tt.position)
		if actual != tt.expected {
			t.Fatal("wrong")
		}
	}
}

func TestUpdateBit(t *testing.T) {
	tests := []struct {
		num      int
		position uint
		bitValue int
		expected int
	}{
		{1, 0, 1, 1},
		{1, 0, 0, 0},
		{1, 1, 1, 3},
		{1, 2, 1, 5},
		{10, 0, 1, 11},
		{10, 0, 0, 10},
		{10, 1, 1, 10},
		{10, 1, 0, 8},
		{10, 2, 1, 14},
		{10, 2, 0, 10},
	}

	for _, tt := range tests {
		actual := UpdateBit(tt.num, tt.position, tt.bitValue)
		if actual != tt.expected {
			t.Fatal("wrong")
		}
	}
}
