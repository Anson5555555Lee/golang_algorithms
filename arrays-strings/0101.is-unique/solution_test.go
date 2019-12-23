package problem0101

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		exp   bool
	}{
		{"valid", "world!", true},
		{"chinese", "你好，世界！", true},
		{"invalid", "Hello, world!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := solution(tt.input)
			if ret != tt.exp {
				t.Fatal("wrong")
			}
		})
	}
}
