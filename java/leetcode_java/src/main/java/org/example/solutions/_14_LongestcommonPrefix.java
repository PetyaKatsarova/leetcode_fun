package org.example.solutions;


/* --- DIFFICULTY EASY ---
* Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".
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
* */

public class _14_LongestcommonPrefix {

    public String longestCommonPrefix(String[] strs) {
        if (strs == null || strs.length == 0)
            return "";
        String  prefix = strs[0];
        int     prefixLength = prefix.length();

        for (int i = 1; i < strs.length; i++) {
            int count = 0;
            int minLen = Math.min(strs[i].length(), prefixLength);
            while ( count < minLen && prefix.charAt(count) == strs[i].charAt(count)) {
                count++;
            }
            prefixLength = count;
        }
        return prefix.substring(0, prefixLength);
    }

    /**
     * in bash from java folder: java org/example/solutions/_14_LongestcommonPrefix.java
     * */

    public static void main(String[] args) {
        System.out.println(new _14_LongestcommonPrefix().longestCommonPrefix(new String[]{"flower", "flow", "flight"})); // fl
        System.out.println(new _14_LongestcommonPrefix().longestCommonPrefix(new String[]{"dog","racecar","car"})); //""
        System.out.println(new _14_LongestcommonPrefix().longestCommonPrefix(new String[]{"a"})); // a
        System.out.println(new _14_LongestcommonPrefix().longestCommonPrefix(new String[]{"aaa","aa","aaa"}));// aa
        System.out.println(new _14_LongestcommonPrefix().longestCommonPrefix(new String[]{"ab", "a"})); // a
        System.out.println(new _14_LongestcommonPrefix().longestCommonPrefix(new String[]{"a","a","b"})); // ""
    }

}