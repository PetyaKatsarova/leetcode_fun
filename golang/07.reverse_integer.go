package main

import (
	"fmt"
	"math"
	// "reflect"
	// "strconv"
)

/* -- TAG -- MEDIUM
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit
 integer range [-2times 31, 2times 31 - 1], then return 0. 
Maximum value for int32: 2,147,483,647 (or 0x7FFFFFFF in hexadecimal)
Minimum value for int32: -2,147,483,648 (or 0x80000000 in hexadecimal)

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21
 
Constraints:
-2 times 31 <= x <= 2times 31 - 1

--theory--
The range of a signed 32-bit integer (int32) in binary is from 
10000000 00000000 00000000 00000000 (which is -2147483648) to
 01111111 11111111 11111111 11111111 (which is 2147483647).
 Bit positions:  31     30     29  ...   2      1      0
Binary value:     1      0      1  ...   0      1      0
The MSB (bit 31) determines the sign of the number.
If the MSB is 0, the number is non-negative (positive or zero).
If the MSB is 1, the number is negative.

Untyped Constants: The literal 1 is an untyped constant. The type
 of 1 << 31 is inferred based on the context where it is used.
Shift Operation: The shift operation << is defined for integer
 types. When you use an untyped constant in a shift expression,
  Go treats the constant as if it has an infinite precision integer
   type, which means it can handle large numbers beyond the typical
    bounds of standard integer types.
Default Type: If an untyped constant is assigned to a variable 
or used in a context where a specific type is expected, Go will
 convert the constant to that type. For example, if you assign
  1 << 31 to an int32 variable, the compiler will ensure
   the value fits within the range of int32.
*/
func main() {
	reverse(123) // 321
	reverse(-123) 
	reverse(120)
	// edge cases: eqaul to max, min and 1 + || -
	reverse(-2147483649) // 0 overflow case
	reverse(2147483648) // 0
	maxInt32 := math.MaxInt32
	minInt32 := math.MinInt32

	fmt.Printf("Max int32: %d\n", maxInt32)
	fmt.Printf("Min int32: %d\n", minInt32)
}

func reverse(x int) int {
	const (
		int32Max = 1<<31 - 1
		int32Min = -1<<31
	)
	var reversed int
	for x != 0 {
		pop := x % 10
		x   = x / 10

		if reversed > int32Max/10 || (reversed / 10 == int32Max && pop > 7) {
			fmt.Println(0)
			return 0
		}
		if reversed < int32Min/10 || (reversed / 10 == int32Min && pop < -8) {
			fmt.Println(0)
			return 0
		}
		reversed = reversed * 10 + pop
	}
	fmt.Println(reversed)
	return reversed
}


// func reverse(x int) int {
// 	const (
// 		// min int32=2times 31    =2147483647
// 		// max int32= -(2times 30)=-2147483648
// 		// move 1 to the left 31 times and remove 1
// 		int32Max = 1<<31 - 1 // -1 for the 0 or the sign or sth
// 		int32Min = -1 << 31
// 	)

// 	var reversed int

// 	for x != 0 {
// 		pop := x % 10
// 		x /= 10

// 		// Check for overflow
// 		if reversed > int32Max/10 || (reversed == int32Max/10 && pop > 7) {
// 			return 0
// 		}
// 		if reversed < int32Min/10 || (reversed == int32Min/10 && pop < -8) {
// 			return 0
// 		}

// 		reversed = reversed*10 + pop
// 	}
// 	fmt.Println(reversed)
// 	return reversed
// }

// func reverse(x int) int {
// 	strResult	:= ""
// 	isMinus		:= false
// 	var intResult int
// 	var err error

// 	if x < 10 && x > -1 {	
// 		fmt.Println("int: ", x)
// 		return x
// 	}

// 	numString := strconv.Itoa(x)
// 	for i := len(numString) - 1; i >= 0; i-- {
// 		strResult += string(numString[i])
// 	}
	
// 	if strResult[len(strResult)-1:] == "-" {
// 		strResult = strResult[:len(strResult)-1]
// 		isMinus = true
// 	}
// 	// do checks for limitations of int32:
// 	if len(strResult) > 2 {
// 		for i, val := range
// 	} else {
// 		if isMinus {
// 			intResult, err = strconv.Atoi(strResult)
// 			if err != nil {panic(err)}
// 			intResult = intResult / -1
// 		} else {
// 			intResult, err = strconv.Atoi(strResult)
// 			if err != nil {panic(err)}
// 		}
// 	}

// 	fmt.Println(intResult, reflect.TypeOf(intResult))

//     return intResult
// }