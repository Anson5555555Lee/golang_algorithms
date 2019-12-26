package tg

import "testing"

import "reflect"

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{4, 2, 1, 3, 6, 5, 7},
		},
	}

	for _, tt := range tests {
		actual := SortedArrayToBST(tt.input).PreOrder()
		if reflect.DeepEqual(actual, tt.expected) != true {
			t.Fatal("wrong")
		}
	}
}
