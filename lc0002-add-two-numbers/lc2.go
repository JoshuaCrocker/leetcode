package lc2addtwonumbers

import "log"

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Constraints:

//     The number of nodes in each linked list is in the range [1, 100].
//     0 <= Node.val <= 9
//     It is guaranteed that the list represents a number that does not have leading zeros.

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(data []int) *ListNode {
	root := &ListNode{}
	ptr := root

	for index, number := range data {
		ptr.Val = number

		if (index + 1) < len(data) {
			ptr.Next = &ListNode{}
			ptr = ptr.Next
		}
	}

	ptr.Next = nil

	return root
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := []int{}

	pointerL1 := l1
	pointerL2 := l2

	overflow := 0

	for pointerL1 != nil || pointerL2 != nil {
		l1Value := 0
		l2Value := 0

		if pointerL1 != nil {
			l1Value = pointerL1.Val
		}

		if pointerL2 != nil {
			l2Value = pointerL2.Val
		}

		total := overflow + l1Value + l2Value

		if total >= 10 {
			overflow = 1
			total = total - 10
		} else {
			overflow = 0
		}

		if pointerL1 != nil {
			pointerL1 = pointerL1.Next
		}
		if pointerL2 != nil {
			pointerL2 = pointerL2.Next
		}

		result = append(result, total)
	}

	if overflow > 0 {
		result = append(result, overflow)
	}

	log.Println(result)
	return CreateLinkedList(result)
}
