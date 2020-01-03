package backtracking

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input    []string
	expected [][]string
}{
	{
		[]string{"a", "b"},
		[][]string{[]string{"a", "b"}, []string{"b", "a"}},
	},
	// {
	// 	[]string{"a", "b", "c"},
	// 	[][]string{
	// 		[]string{"a", "b", "c"},
	// 		[]string{"a", "c", "b"},
	// 		[]string{"b", "a", "c"},
	// 		[]string{"b", "c", "a"},
	// 		[]string{"c", "a", "b"},
	// 		[]string{"c", "b", "a"},
	// 	},
	// },
}

func TestPermutate(t *testing.T) {
	for _, tt := range tests {
		actual := permutate(tt.input)
		if reflect.DeepEqual(actual, tt.expected) != true {
			t.Fatal("wrong")
		}
	}
}
