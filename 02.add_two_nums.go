package main

/* --- TAG: MEDIUM ----
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
 and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:
The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros
*/

import (
	"fmt"
)

func main() {
	fmt.Println("hello add 2 nums linked lists")
	l1 		:= &ListNode{2, &ListNode{4, &ListNode{3, nil}}} // 243
	l2 		:= &ListNode{5, &ListNode{6, &ListNode{4, nil}}} // 564
// 	Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]
	l3		:= &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l4		:= &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	result	:= addTwoNumbers(l1, l2) // [7,0,8]
	//7->0->8
	for result != nil {
		fmt.Print(result.Val, ",")
		result = result.Next
	}
	result2	:= addTwoNumbers(l3, l4) // [8,9,9,9,0,0,0,1]
	fmt.Println("")
	for result2 != nil {
		fmt.Print(result2.Val, ",")
		result2 = result2.Next
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers (l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead	:= &ListNode{0, nil}
	curr 		:= dummyHead
	carryOver	:= 0

	for l1 != nil || l2 != nil || carryOver != 0 {
		sum := carryOver
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carryOver = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
	}
	return dummyHead.Next
}
