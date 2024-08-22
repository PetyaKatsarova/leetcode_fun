package main

/* -- TAG -- HARD
https://leetcode.com/problems/trapping-rain-water/description/
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1] shows only the y coord in the coordinate system, x is the index of the slice
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
Example 2:
Input: height = [4,2,0,3,2,5]
Output: 9
restrictions:
n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/

// import "fmt"
// func main() {
// 	fmt.Println("hello world")
// 	trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}) // 6
// 	trap([]int{4,2,0,3,2,5}) // 9
// }

// func trap(height []int) int {
// 	water := 0
// 	left, right := 0, len(height)-1
// 	maxLeft, maxRight := 0, 0

// 	for left < right {
// 		if height[left] < height[right] {
// 			if maxLeft < height[left] {
// 				maxLeft = height[left]
// 			} else {
// 				water += maxLeft - height[left]
// 			}
// 			left++
// 		} else {
// 			if maxRight < height[right] {
// 				maxRight = height[right]
// 			} else {
// 				water += maxRight - height[right]
// 			}
// 			right--
// 		}
// 	}

// 	fmt.Println("water is: ", water)
// 	return water;
// }

