package org.example.solutions;


import java.util.Map;
import static java.util.Map.entry;
/**
 * run in src/main/java:
 *  javac org/example/solutions/_13_Roman_to_int.java
 * java org.example.solutions._13_Roman_to_int
 * if u want to print varargs:  java org.example.solutions._13_Roman_to_int bla bla bla
 * FROM JAVA 11, NO NEED TO COMPILE, DO ONLY:
 *  java org/example/solutions/_13_Roman_to_int.java kuku lala
 * */

/* -- DIFFICULTY TAG -- EASY
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as
 XXVII, which is XX + V + II. Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
 Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine,
  which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.

Example 1:
Input: s = "III"
Output: 3
Explanation: III = 3.

Example 2:
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Example 3:
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

Constraints:
1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999]. MMM=3000; CM=900; XC=90; IX=9; 3999=MMMCMXCIX;

        System.out.println(table.get("I"));  // Output: 1
        System.out.println(table.get("V"));  // Output: 5
        *  Map<String, Integer> table = Map.of(
            "I",1,
            "IV", 4,
            "V", 5,
            "IX", 9,
            "X", 10
        );
    }
}
Map.ofEntries() allows you to add more than 10 entries (which is the limit for Map.of()), and uses Map.entry() to add the entries in a readable manner.

NB: Runtime 4 ms Beats 67.30% Memory 44.31 MB Beats 91.55%
*/
public class _13_RomanToInt {
    private final static Map<String, Integer> TABLE = Map.ofEntries(
            entry("I", 1),
            entry("IV", 4),
            entry("V", 5),
            entry("IX", 9),
            entry("X", 10),
            entry("XL", 40),
            entry("L", 50),
            entry("XC", 90),
            entry("C", 100),
            entry("CD", 400),
            entry("D", 500),
            entry("CM", 900),
            entry("M", 1000)
    );

    public static int convertRom2Int(String roman) {
       int  result = 0;
       int  counter = 0;
       while(counter < roman.length()) {
           if (counter + 1 < roman.length()&& TABLE.containsKey(roman.substring(counter, counter + 2))) {
               result += TABLE.get(roman.substring(counter, counter + 2));
               counter += 2;
           } else {
               result += TABLE.get(roman.substring(counter, counter+1));
               counter++;
           }
       }
        return result;
    }
    public static void main(String[] args) {
        System.out.println("MCMXCIV(1994) = " + convertRom2Int("MCMXCIV")); // 1994
        System.out.println("LVIII(58)= " + convertRom2Int("LVIII")); // 58
//        System.out.println(args[0] + ", " + args[1]); // for this to compile add args
    }
}
