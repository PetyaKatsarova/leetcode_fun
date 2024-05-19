/*
Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.
The algorithm for myAtoi(string s) is as follows:
Whitespace: Ignore any leading whitespace (" ").
Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity is neither present.
Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read,
 then the result is 0.
 Rounding: If the integer is out of the 32-bit signed integer range [-2times31, 2times31 - 1], then round the integer to remain in the range.
 Specifically, integers less than -2times31 should be rounded to -2times31, and integers greater than 2times31 - 1 should be rounded to 2times31 - 1.
Return the integer as the final result.

Example 1:
Input: s = "42"
Output: 42

Explanation:
The underlined characters are what is read in and the caret is the current reader position.
Step 1: "42" (no characters read because there is no leading whitespace)
         ^
Step 2: "42" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "42" ("42" is read in)
           ^
Example 2:
Input: s = " -042"
Output: -42

Explanation:
Step 1: "   -042" (leading whitespace is read and ignored)
            ^
Step 2: "   -042" ('-' is read, so the result should be negative)
             ^
Step 3: "   -042" ("042" is read in, leading zeros ignored in the result)
               ^
Example 3:
Input: s = "1337c0d3"
Output: 1337

Explanation:
Step 1: "1337c0d3" (no characters read because there is no leading whitespace)
         ^
Step 2: "1337c0d3" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "1337c0d3" ("1337" is read in; reading stops because the next character is a non-digit)
             ^
Example 4:
Input: s = "0-1"
Output: 0

Explanation:
Step 1: "0-1" (no characters read because there is no leading whitespace)
         ^
Step 2: "0-1" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "0-1" ("0" is read in; reading stops because the next character is a non-digit)
          ^
Example 5:

Input: s = "words and 987"
Output: 0

Explanation:
Reading stops at the first non-digit character 'w'.

Constraints: 0 <= s.length <= 200
s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.
*/

package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func main() {
	fmt.Println("expected: 0; output: ", myAtoi("  -00kuku 0042v7"))
	fmt.Println("expected: -42, output: ", myAtoi("  -042"))
    fmt.Println("expected: 2147483646, output: ", myAtoi("2147483646"))
	fmt.Println("expected: -2147483648, output: ", myAtoi("-91283472332"))
	fmt.Println("expected: -2147483648, output: ", myAtoi(" -2147483649"))
}

func myAtoi(s string) int {
    if len(s) == 0 { return 0 }
	isNeg := false
	s = ignoreLeadingSpace(s)
    if s == "" { return 0 }
	if s[0] == '-' { // determine the sign: works
		isNeg = true
		s = s[1:]
	} else if s[0] == '+' {
        s = s[1:]
    }
	i := 0
	for i < len(s) && s[i] == '0' { //Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read,then the result is 0.
		i++
	}
	if i == len(s) { // if there is no num in the string return 0
		return 0
	} else { s = s[i:] }
	re          := regexp.MustCompile(`^\d+`) // starts with digit one or many digits; 	// encounter non digit character:
	s           = re.FindString(s)
	return numFromStr(s, isNeg)
}

func ignoreLeadingSpace(s string) string {
    i := 0
    for i < len(s) && unicode.IsSpace(rune(s[i])) { //1. ignore any leading whitespace: works
		i++
	}
    if i == len(s) {
        return ""
    }
	return s[i:]
}

func numFromStr(s string, isNeg bool) int {
    max32Int    := 1<<31 - 1 // 2147483647
	min32Int    := -1 << 31  // -2147483648
	result      := 0

	for i := range s { // convert string of digits to num
		digit := int(s[i] - '0')

		if result*10+digit > max32Int || result*10+digit < min32Int {
			if isNeg {
				return min32Int
			}
			return max32Int
		}
		result = result*10 + digit
	}
	if isNeg {
		return -result
	}
    return result
}
