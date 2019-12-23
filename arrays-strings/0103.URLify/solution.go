package problem0103

import "fmt"

// Time complexity O(n), space complexity O(n)
func solution1(input string) string {
	count, spaces := 0, 0

	for _, v := range input {
		count++
		if v == rune(' ') {
			spaces++
		}
	}

	temp := make([]rune, count+2*spaces)
	ptr := 0
	for _, v := range input {
		if v == rune(' ') {
			temp[ptr] = rune('%')
			temp[ptr+1] = rune('2')
			temp[ptr+2] = rune('0')
			ptr += 3
		} else {
			temp[ptr] = v
			ptr++
		}
	}

	return string(temp)
}

// How to slove this problem without extra spaces appended???

// In-place implementation
func URLifySlice(input []rune) {
	inWord := false
	slowPtr := len(input) - 1
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == rune(' ') {
			if inWord {
				input[slowPtr-2] = rune('%')
				input[slowPtr-1] = rune('2')
				input[slowPtr] = rune('0')
				slowPtr -= 3
			}
		} else {
			if !inWord {
				inWord = true
			}
			input[slowPtr] = input[i]
			slowPtr--
		}

		s := fmt.Sprintf("%s", string(input))
		fmt.Println(s)
	}
}
