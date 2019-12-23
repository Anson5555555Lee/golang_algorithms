package problem0102

import "testing"

func TestSolution1(t *testing.T) {
	tests := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"abcd", "abcd", true},
		{"abcd", "abdc", true},
		{"abcc", "ccbb", false},
		{"abcc", "abcc ", false},
		{" ", " ", true},
		{"", "", true},
	}

	for _, tt := range tests {
		actual := solution1(tt.input1, tt.input2)
		if tt.expected != actual {
			t.Fatalf("Inputs %s, %s. Expected: %t, actual: %t\n",
				tt.input1, tt.input2, tt.expected, actual)
		}
	}
}

func TestSolution2(t *testing.T) {
	tests := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"abcd", "abcd", true},
		{"abcd", "abdc", true},
		{"abcc", "ccbb", false},
		{"abcc", "abcc ", false},
		{" ", " ", true},
		{"", "", true},
	}

	for _, tt := range tests {
		actual := solution2(tt.input1, tt.input2)
		if tt.expected != actual {
			t.Fatalf("Inputs %s, %s. Expected: %t, actual: %t\n",
				tt.input1, tt.input2, tt.expected, actual)
		}
	}
}
