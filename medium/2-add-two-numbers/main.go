package main

// https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	current := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		sum %= 10
		current.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
		current = current.Next
	}

	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	sum := decimalValue(l1) + decimalValue(l2)
//	node := newListNode(sum)
//	return &node
//}
//
// func decimalValue(l *ListNode) int {
// 	sum := 0
// 	iter := 1
// 	node := l
// 	for node != nil {
// 		sum += node.Val * iter
// 		iter *= 10
// 		node = node.Next
// 	}
// 	return sum
// }
//
// func newListNode(decimal int) ListNode {
// 	var val int
// 	var arr []ListNode
// 	for {
// 		val = decimal % 10
// 		decimal /= 10
// 		arr = append(arr, ListNode{
// 			Val:  val,
// 			Next: nil,
// 		})
// 		if decimal == 0 {
// 			break
// 		}
// 	}
//
// 	for i := 0; i < len(arr)-1; i++ {
// 		arr[i].Next = &arr[i+1]
// 	}
//
// 	return arr[0]
// }
//
