package main

import (
	"fmt"
)

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains
 a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
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
It is guaranteed that the list represents a number that does not have leading zeros.

2->4->3
5->6->4
----------------
7->0->8
*/

func main() {
	fmt.Println("hello add 2 nums linked lists")
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	//7->0->8
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead	:= &ListNode{0, nil}
	current		:= dummyHead
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
		current.Next = &ListNode{sum % 10, nil}
		current = current.Next
	}
	return dummyHead.Next
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	sum := convertListToReversedNum(l1) + convertListToReversedNum(l2)
// 	fmt.Println("sum is: ", sum)
// 	result := ListNode{}
// 	sumStr := strconv.Itoa(sum)
// 	//l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
// 	// reverse int in the num:
// 	for i, j := 0, len(sumStr); i < len(sumStr), j > 0; i++, j-- {

// 	}
// 	return &result
// }

// func convertListToReversedNum(list *ListNode) int {
// 	node1	:= list
// 	sl		:= []int{}

// 	for node1 != nil { 
// 		sl = append(sl, node1.Val)
// 		node1 = node1.Next
// 	}

// 	n 			:= 0
// 	multiplyBy 	:= 1
// 	for i := 0; i < len(sl); i++ {
// 		n += sl[i] * multiplyBy;
// 		multiplyBy *= 10
// 	}
	
// 	fmt.Println("num------------", n) //342
// 	return n
// }
// //2->4->3
// // 5->6->4
// // ----------------
// // 7->0->8
