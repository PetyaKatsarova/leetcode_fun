package main

/* -- TAG -- medium
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith
 line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.
Example 1: Input: height = [1,8,6,2,5,4,8,3,7] Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can 
//contain is 49.
Example 2: Input: height = [1,1] Output: 1
Constraints:
n == height.length
2 <= n <= 105
0 <= height[i] <= 10000
*/

import "fmt"
func main() {
	fmt.Println("hallo world")
	fmt.Println("expected: 49, outcome: ", maxArea([]int{1,8,6,2,5,4,8,3,7})) // 49
	fmt.Println("expected: 1, outcome: ", maxArea([]int{1,1})) //1
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea 	:= 0

	for left < right {
		minHeight := height[left]
		if height[left] > height[right] {
			minHeight = height[right]
		}
		area := minHeight * (right - left)
		if area > maxArea { maxArea = area }
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// too slow: O(n^2) time complexity: nested loops
// func maxArea(height []int) int {
//     maxArea := 0
// 	for i:=0; i<len(height); i++ {
// 		for j:=0; j<len(height); j++ {
// 			minHeight := height[i] 
// 			if minHeight > height[j] { minHeight = height[j] }

// 			if maxArea < minHeight * (j - i) { maxArea = minHeight * (j - i) }
// 		}
// 	}
// 	return maxArea
// }
