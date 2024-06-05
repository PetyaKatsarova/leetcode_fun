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
	lenS, lenP := len(s), len(p)
	twoD := make([][]bool, lenS+1) // default vals = false; if s=="abscd": [[false, false, false, false],[],[],[],[]]

	for i := range twoD {
		twoD[i] = make([]bool, lenP+1) /* if p=="a*d": [
			[false, false, false, false],
			[false, false, false, false],
			[false, false, false, false],
			[false, false, false, false],
			[false, false, false, false],
		]

		*/
	}
	twoD[0][0] = true
	for j := 1; j <= lenP; j++ {
		if p[j-1] == '*' {
			twoD[0][j] = twoD[0][j-2]
		}
	}
}

// func isMatch(s string, p string) bool {
// 	m, n := len(s), len(p)
// 	dp := make([][]bool, m+1)
// 	for i := range dp {
// 		dp[i] = make([]bool, n+1)
// 	}
// 	dp[0][0] = true

// 	for j := 1; j <= n; j++ {
// 		if p[j-1] == '*' {
// 			dp[0][j] = dp[0][j-2]
// 		}
// 	}

// 	for i := 1; i <= m; i++ {
// 		for j := 1; j <= n; j++ {
// 			if p[j-1] == '.' || p[j-1] == s[i-1] {
// 				dp[i][j] = dp[i-1][j-1]
// 			} else if p[j-1] == '*' {
// 				dp[i][j] = dp[i][j-2]
// 				if p[j-2] == '.' || p[j-2] == s[i-1] {
// 					dp[i][j] = dp[i][j] || dp[i-1][j]
// 				}
// 			}
// 		}
// 	}
// 	return dp[m][n]
// }

