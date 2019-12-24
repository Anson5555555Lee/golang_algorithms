package problem0104

func solution(input string) bool {
	counts := make(map[rune]int)
	len := len(input)
	for _, r := range input {
		counts[r]++
	}

	odd := len%2 == 1
	for _, v := range counts {
		if v%2 == 1 {
			if odd == false {
				return false
			}
		}
	}

	return true
}
