package problem0102

import (
	"sort"
	"strings"
)

// Solution #1: Sort the strings.
// In O(N*logN) time complexity
func solution1(input1, input2 string) bool {
	if len(input1) != len(input2) {
		return false
	}

	if sortHelper(input1) == sortHelper(input2) {
		return true
	}

	return false
}

// Solution #2: Check if the strings have identical character counts.
// In O(N) time complexity
func solution2(input1, input2 string) bool {
	if len(input1) != len(input2) {
		return false
	}

	counts := make(map[rune]int)

	for _, r := range input1 {
		counts[r]++
	}

	for _, r := range input2 {
		counts[r]--
	}

	for _, val := range counts {
		if val != 0 {
			return false
		}
	}

	return true
}

// What is the efficiency of sorting algorithm in golang
func sortHelper(input string) string {
	s := strings.Split(input, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
