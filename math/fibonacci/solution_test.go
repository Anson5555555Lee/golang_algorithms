package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		data int
		exp  []int
	}{
		{
			"1",
			1,
			[]int{1},
		},
		{
			"2",
			2,
			[]int{1, 1},
		},
		{
			"3",
			3,
			[]int{1, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := fibonacci(tt.data)
			if reflect.DeepEqual(ret, tt.exp) != true {
				t.Fatal("wrong")
			}
		})
	}
}

func TestFibonacciNth(t *testing.T) {
	tests := []struct {
		name  string
		index int
		exp   int
	}{
		// {"0", 0, 0},
		// {"1", 1, 1},
		// {"2", 2, 1},
		// {"3", 3, 2},
		// {"4", 4, 3},
		// {"5", 5, 5},
		// {"6", 6, 8},
		{"7", 7, 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := fibonacciNth(tt.index, make([]int, tt.index+1))
			if reflect.DeepEqual(ret, tt.exp) == false {
				t.Fatal("wrong")
			}
		})
	}
}
