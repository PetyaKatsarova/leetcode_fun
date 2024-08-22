package org.example.solutions;

import java.util.Map;

public class Solution_13 {

    private final static Map<Character, Integer> TABLE = Map.of(
            'I', 1,
            'V', 5,
            'X', 10,
            'L', 50,
            'C', 100,
            'D', 500,
            'M', 1000
    );

    public static int romanToInt(String s) {
        int result = 0;
        int length = s.length();

        for (int i = 0; i < length; i++) {
            int currentValue = TABLE.get(s.charAt(i));
            if (i + 1 < length && currentValue < TABLE.get(s.charAt(i + 1))) {
                result -= currentValue;
            } else {
                result += currentValue;
            }
        }

        return result;
    }
    public static void main(String[] args) {
        System.out.println("MCMXCIV(1994) = " + romanToInt("MCMXCIV")); // 1994
        System.out.println("LVIII(58)= " + romanToInt("LVIII")); // 58
//        System.out.println(args[0] + ", " + args[1]); // for this to compile add args
    }
}
