package main

import "fmt"

/*
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
	fmt.Println("hello two sums")
	twoSum([]int{2,7,11,15}, 9) // [0,1]
	twoSum([]int{3,2,4}, 6) // [1,2]
	twoSum([]int{3,3}, 6) // [0,1]
}

func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int)
	for i, val := range nums {
		remainder := target - val;
		if j, ok := mapNums[remainder]; ok {
			fmt.Println("result is: ", []int{j,i})
            return []int{j,i}
		}
		mapNums[val] = i;
	}
	return []int{-1,-1}
}

// func twoSum(nums []int, target int) []int {
//    for i := 0; i< len(nums) - 1; i++ {
// 		for j := 1; j < len(nums); j++ {
// 			if nums[i] + nums[j] == target {
// 				fmt.Println("result: ", []int{i, j})
// 				return []int{i, j}
// 			}
// 		}
//    }
// 	return []int{-1, -1}
// }