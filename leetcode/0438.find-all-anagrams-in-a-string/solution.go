package problem0438

func findAnagrams(s string, p string) []int {

	M, N := len(p), len(s)

	if M > N {
		return []int{}
	}

	result := []int{}

	var target, window [26]int
	for i := 0; i < M; i++ {
		target[p[i]-'a']++
		window[s[i]-'a']++
	}

	for i := M; i < N; i++ {
		if window == target {
			result = append(result, i-M)
		}

		window[s[i]-'a']++
		window[s[i-M]-'a']--
	}

	if window == target {
		result = append(result, N-M)
	}

	return result
}
