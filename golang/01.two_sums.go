package main

import "fmt"

/* -- TAG: EASY  ---
https://leetcode.com/problems/trapping-rain-water/description/
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
Constraints:
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
*/

func main() {
	fmt.Println("hellow 2 sums :)")
	twosums([]int{2,7,11,15}, 9) // [0,1]
	twosums([]int{3,2,4}, 6) // [1,2]
}


func twosums(nums []int, n int) (result []int) {
	mymap := make(map[int]int)
	for i, val := range nums {
		remainder := n - val
		if j, ok := mymap[remainder]; ok {
			fmt.Println("result: ", j, i)
			return []int{j, i}
		}
		mymap[val] = i
	}
	fmt.Println(result)
	return result
}

