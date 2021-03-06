package quick

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		data []int
		exp  []int
	}{
		{
			"1",
			[]int{1},
			[]int{1},
		},
		{
			"12",
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			"21",
			[]int{2, 1},
			[]int{1, 2},
		},
		{
			"reversed",
			[]int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
		{
			"notSorted",
			[]int{15, 8, 5, 12, 10, 1, 16, 9, 11, 7, 20, 3, 2, 6, 17, 18, 4, 13, 14, 19},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
		{
			"negative",
			[]int{-1, 0, 5, -10, 20, 13, -7, 3, 2, -3},
			[]int{-10, -7, -3, -1, 0, 2, 3, 5, 13, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sort(tt.data)
			if ret := reflect.DeepEqual(res, tt.exp); ret != true {
				t.Fatalf("wrong")
			}
		})
	}
}
