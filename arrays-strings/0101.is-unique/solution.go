package problem0101

// Using map of runes for duplicate detection.
func solution(data string) bool {
	seen := make(map[rune]struct{})
	for _, v := range data {
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = struct{}{}
	}

	return true
}
