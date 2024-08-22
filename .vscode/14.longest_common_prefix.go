package main

import (
	"fmt"
)

/*  ----- TAG ---- EASY -------
Write a function to find the longest common prefix string amongst an array of strings. If there is no common prefix, return an empty string "".
Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
*/

func main() {
	fmt.Println("expected: 'fl', output: ", longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println("expected: '', output: ", longestCommonPrefix([]string{"dog","docar","car"}))
}

func longestCommonPrefix(strs []string) string { 
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for j < min(len(strs[i]), len(prefix)) && prefix[j] == strs[i][j] {
			j++
		}
		prefix = strs[i][:j]
		if prefix == "" { return ""}
	}
	return prefix
}
