package main

import "fmt"

/* --- TAG: HARD -----------------
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).

Example 1:
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

Example 2:
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

Constraints:
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106 // can use int16
*/

func main() {
	findMedianSortedArrays([]int{1, 2}, []int{3, 4}) // 2.5
	 findMedianSortedArrays([]int{1, 3}, []int{2})    // 2
	fmt.Println(mergeArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(mergeArrays([]int{1, 3}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	median := 0.0
	mergedArr := mergeArrays(nums1, nums2)
	if len(mergedArr) % 2 == 0 { // even num
		median = (float64(mergedArr[len(mergedArr) / 2]) + float64(mergedArr[(len(mergedArr) / 2) - 1])) / 2
	} else { // odd num
		median = float64(mergedArr[(len(mergedArr) / 2)])
	}

	fmt.Println("the median of the 2 arr is: ", median)
	return median
}

func mergeArrays(nums1 []int, nums2 []int) []int {
	mergedArr	:= make([]int, len(nums1)+len(nums2))
	i, j, k		:= 0, 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			mergedArr[k] = nums2[j]
			j++
		} else {
			mergedArr[k] = nums1[i]
			i++
		}
		k++
	}
	for i < len(nums1) {
		mergedArr[k] = nums1[i]
		k++
		i++
	}
	for j < len(nums2) {
		mergedArr[k] = nums2[j]
		k++
		j++
	}
	return mergedArr
}
