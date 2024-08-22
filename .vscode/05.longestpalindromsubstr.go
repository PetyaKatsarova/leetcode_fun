package main

import "fmt"

/*
 --- TAG --- MEDIUM
 Given a string s, return the longest
palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"

Constraints:
1 <= s.length <= 1000
s consist of only digits and English letters.
*/

func main() {
	longestPalindrome("ccc") // ccc
	longestPalindrome("babad") // bab
	longestPalindrome("cbbd") // bb
}

// from a string get a start end end of polindrom
func getPalindrom(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := range s {
		leftOdd, rightOdd := getPalindrom(s, i, i) // for odd len of the palindrom
		leftEven, rightEven := getPalindrom(s, i, i + 1)

		if rightOdd - leftOdd > end - start {
			start = leftOdd
			end   = rightOdd
		}
		if rightEven - leftEven > end - start {
			start = leftEven
			end   = rightEven
		}
	}
	fmt.Println(s[start:end+1])
	return s[start:end+1]
}
