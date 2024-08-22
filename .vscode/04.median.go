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
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))   // 2
	fmt.Println(mergeArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(mergeArrays([]int{1, 3}, []int{2}))
}

func mergeArrays(n1 []int, n2 []int) []int {
	myArr 	:= make([]int, len(n1)+len(n2))
	i, j, k	:= 0, 0, 0

	for i < len(n1) && j < len(n2) {
		if n1[i] < n2[j] {
			myArr[k] = n1[i]
			i++
		} else {
			myArr[k] = n2[j]
			j++
		}
		k++
	}

	for i < len(n1) {
		myArr[k] = n1[i]
		i++
		k++
	}

	for j < len(n2) {
		myArr[k] = n2[j]
		j++
		k++
	}

	return myArr
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArr  := mergeArrays(nums1, nums2)

	if len(mergedArr) % 2 == 0 {
		return (float64(mergedArr[len(mergedArr) / 2]) + float64(mergedArr[(len(mergedArr) / 2) - 1])) / 2
	} else {
		return float64(mergedArr[len(mergedArr) / 2]) // if len=5, 5/2 = 2 (index 2, 3rd middle num)
	}
}
