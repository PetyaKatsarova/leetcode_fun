/* -- TAG -- HARD
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Example 1:
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Example 2:
Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

Example 3:
Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".

Constraints:
1 <= s.length <= 20
1 <= p.length <= 20
s contains only lowercase English letters.
p contains only lowercase English letters, '.', and '*'.
It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("expected: false(for p 'a'), output: ", isMatch("aa", "a"))
	// fmt.Println("expected: false(for p 'aaa'), output: ", isMatch("aa", "aaa"))
	fmt.Println("expected: true(for p 'a*'), output: ", isMatch("aa", "a*"))
	// fmt.Println("expected: true(for p '.*'), output: ", isMatch("abc", "."))
	// fmt.Println("expected: true(for p 'a.*'), output: ", isMatch("ab", "a."))
	// fmt.Println("expected: false(for p 'a.*'), output: ", isMatch("abc", "a."))
	// fmt.Println("expected: true(for p 'c4a2b'), output: ", isMatch("c4a2b", "c.a.b"))
	// fmt.Println("expected: true(for p 'c4a42b'), output: ", isMatch("c4a42b", "c.a..b"))
	// fmt.Println("expected: false(for p 'c4a4b'), output: ", isMatch("c4a4b", "c.a..b"))
	// fmt.Println("expected: true(for p 'c4a42b...'), output: ", isMatch("c4a42bbla", "c.a..b..."))
}

func isMatch(s string, p string) bool {
	// 1. no . and no *
	if !strings.Contains(p, "*") && !strings.Contains(p, ".") {
		return noSpecialChars(s, p)
	}

	// 2. only . (wild card)
	if !strings.Contains(p, "*") && strings.Contains(p, ".") {
		return onlyDot(s, p)
	}
	return false
}

func noSpecialChars(s string, p string) bool {
	if len(s) != len(p) {
		return false
	}
	i := 0
	for i < len(s) {
		if s[i] != p[i] {
			return false
		}
		i++
	}
	return true
}

func onlyDot(s string, p string) bool {
	if len(s) != len(p) {
		return false
	}
	for i := range s {
		if s[i] != p[i] && p[i] != '.' {
			fmt.Println("len(s) & len(p): ", len(s), len(p), string(s[i]), string(p[i])) // isMatch("c4a4b", "c.a..b...")
			return false
		}
	}
	return true
}

func starAndDot(s string, p string) bool { // isMatch("aa", "a*")
	j := 0
	for i := range s {
		
	}
	return true
}

// 1. no . and no *: only chars
// 2. only .
// 3. * and .

