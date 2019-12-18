package problem0009

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		data int
		exp  bool
	}{
		{
			"negative",
			-12321,
			false,
		},
		{
			"zero",
			0,
			true,
		},
		{
			"invalid",
			123,
			false,
		},
		{
			"valid",
			12321,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome(tt.data)
			if ret := reflect.DeepEqual(res, tt.exp); ret != true {
				t.Fatalf("wrong")
			}
		})
	}
}
