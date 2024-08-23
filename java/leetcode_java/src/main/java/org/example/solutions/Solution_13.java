package org.example.solutions;

import java.util.Map;
import static java.util.Map.entry;

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

        for (int i = 0; i < s.length(); i++) {
            int curr = TABLE.get(s.charAt(i));
            if (i + 1 < s.length() && curr < TABLE.get(s.charAt(i+1)))
                result -= TABLE.get(s.charAt(i));
            else
                result += TABLE.get(s.charAt(i));
        }
        return result;
    }
    public static void main(String[] args) {
        System.out.println("MCMXCIV(1994) = " + romanToInt("MCMXCIV")); // 1994
        System.out.println("LVIII(58)= " + romanToInt("LVIII")); // 58
        System.out.println("III(3)= " + romanToInt("III")); // 3
//        System.out.println(args[0] + ", " + args[1]); // for this to compile add args
    }
}
