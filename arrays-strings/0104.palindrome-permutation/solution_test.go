package problem0104

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"amanaplanacanalpanama", true},   // normal palindrome
		{"amanalpancaaanplanama", true},   // jumbled palindrome
		{"amanaplanacanalpanamab", false}, // not a palindrome
		{"a", true},
		{"", true},
	}
	for _, c := range cases {
		actual := solution(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %t, actual: %t\n", c.input, c.expected, actual)
		}
	}
}
