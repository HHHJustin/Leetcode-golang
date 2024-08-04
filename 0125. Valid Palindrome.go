package main

import "unicode"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		l := rune(s[left])
		r := rune(s[right])
		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left++
		} else if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			right--
		} else if unicode.ToLower(l) == unicode.ToLower(r) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

// s = [r, a, c, e, , a, c, a, r]
//      ^ i                    ^ j -> if s[i] is letter
//                                 -> if s[j] is letter -> tolower and compare
// s[i] == s[j] -> i++, j-- while i <= j
// ...
// s = [r, a, c, e, , a, c, a, r]
//               ^ i  ^ j -> if s[i] is letter
//                      -> if s[j] is letter -> tolower and compare
// s[i] != s[j] -> return false

// unicode.IsLetter -> check if letter.
// unicode.IsDigit -> check if digit.
// unicode.ToLower -> transfer letter to lower.
