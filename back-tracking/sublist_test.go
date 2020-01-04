package backtracking

import (
	"reflect"
	"testing"
)

var testsSublists = []struct {
	input    []string
	expected [][]string
}{
	{
		[]string{"Jane", "Bob", "Matt", "Sara"},
		[][]string{
			[]string{"Jane", "Bob", "Matt", "Sara"},
			[]string{"Jane", "Bob", "Matt"},
			[]string{"Jane", "Bob", "Sara"},
			[]string{"Jane", "Bob"},
			[]string{"Jane", "Matt", "Sara"},
			[]string{"Jane", "Matt"},
			[]string{"Jane", "Sara"},
			[]string{"Jane"},
			[]string{"Bob", "Matt", "Sara"},
			[]string{"Bob", "Matt"},
			[]string{"Bob", "Sara"},
			[]string{"Bob"},
			[]string{"Matt", "Sara"},
			[]string{"Matt"},
			[]string{"Sara"},
			[]string{},
		},
	},
	{
		[]string{"a", "b"},
		[][]string{
			[]string{},
			[]string{"b"},
			[]string{"a"},
			[]string{"a", "b"},
		},
	},
}

func TestSublists(t *testing.T) {
	for i, tt := range testsSublists {
		actual := sublists(tt.input)
		if i == 0 {
			for i, j := 0, len(actual)-1; i < j; i, j = i+1, j-1 {
				actual[i], actual[j] = actual[j], actual[i]
			}
		}
		if reflect.DeepEqual(actual, tt.expected) != true {
			t.Fatal("wrong")
		}
	}
}
