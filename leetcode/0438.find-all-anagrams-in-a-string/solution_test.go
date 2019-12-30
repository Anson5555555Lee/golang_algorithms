package problem0438

import "testing"

import "reflect"

var tests = []struct {
	s        string
	p        string
	expected []int
}{
	{"abc", "cbaebabacd", []int{}},
	{"cbaebabacd", "abc", []int{0, 6}},
	{"abab", "ab", []int{0, 1, 2}},
	{"", "a", []int{}},
}

func TestFindAnagrams(t *testing.T) {
	for _, tt := range tests {
		actual := findAnagrams(tt.s, tt.p)
		if reflect.DeepEqual(actual, tt.expected) == false {
			t.Fatal("wrong")
		}
	}
}
