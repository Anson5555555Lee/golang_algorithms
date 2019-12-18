package problem0009

import "strconv"

// a solution without converting the integer to a string
// func isPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}

// 	reversedX := 0
// 	originalX := x
// 	for x > 0 {
// 		reversedX = reversedX*10 + x%10
// 		x /= 10
// 	}

// 	if reversedX == originalX {
// 		return true
// 	}

// 	return false
// }

// a solution without extra memory space
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
