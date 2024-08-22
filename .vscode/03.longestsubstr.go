package main

/*
-- TAG: MEDIUM  -----
Given a string s, find the length of the longest substring without repeating characters.
Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

import (
	"fmt"
)

func main() {
	fmt.Println("hello from 03 :)")
	// fmt.Println("output of abcabcbb, expected 3, actual: ", lengthOfLongestSubstring("abcabcbb"))
	// fmt.Println("output of bbbbb, expected 1, actual: ", lengthOfLongestSubstring("bbbbb"))
	// fmt.Println("output of pwwkew, expected 3, actual: ", lengthOfLongestSubstring("pwwkew"))
	fmt.Println("output of dvdf, expected 3, actual: ", lengthOfLongestSubstring("dvdf")) // vdf
}

func lengthOfLongestSubstring(s string) int {
	mymap	:= make(map[rune]int)
	maxLen 	:= 0
	start	:= 0
	for i, val := range s {
		if j, ok := mymap[val]; ok && j >= start {
			start = j + 1
		}
		mymap[val] = i
		if curLen := i - start + 1; curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
