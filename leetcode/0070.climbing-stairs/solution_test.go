package problem0070

import "testing"

func TestBruteForce(t *testing.T) {
	tests := []struct {
		name string
		data int
		exp  int
	}{
		{
			"1 steps",
			1,
			1,
		},
		{
			"2 steps",
			2,
			2,
		},
		{
			"3 steps",
			3,
			3,
		},
		{
			"4 steps",
			4,
			5,
		},
		{
			"5 steps",
			5,
			8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := bruteForce(tt.data)
			if ret != tt.exp {
				t.Fatal("wrong")
			}
		})
	}
}

func TestMemoization(t *testing.T) {
	tests := []struct {
		name string
		data int
		exp  int
	}{
		{
			"1 steps",
			1,
			1,
		},
		{
			"2 steps",
			2,
			2,
		},
		{
			"3 steps",
			3,
			3,
		},
		{
			"4 steps",
			4,
			5,
		},
		{
			"5 steps",
			5,
			8,
		},
		{
			"6 steps",
			6,
			13,
		},
		{
			"7 steps",
			7,
			21,
		},
		{
			"8 steps",
			8,
			34,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := memoization(tt.data, make(map[int]int, tt.data+1))
			if ret != tt.exp {
				t.Fatal("wrong")
			}
		})
	}
}

func TestTripleSteps(t *testing.T) {
	tests := []struct {
		name string
		data int
		exp  int
	}{
		{
			"1 steps",
			1,
			1,
		},
		{
			"2 steps",
			2,
			2,
		},
		{
			"3 steps",
			3,
			4,
		},
		{
			"4 steps",
			4,
			7,
		},
		{
			"5 steps",
			5,
			13,
		},
		{
			"6 steps",
			6,
			24,
		},
		{
			"7 steps",
			7,
			44,
		},
		{
			"8 steps",
			8,
			81,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := tripleSteps(tt.data, make(map[int]int, tt.data+1))
			if ret != tt.exp {
				t.Fatal("wrong")
			}
		})
	}
}
