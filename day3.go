package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val int
	Next *ListNode
 }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, tail, cur *ListNode
	p1, p2 := l1, l2

	for {
		if p1 == nil && p2 == nil {
			break
		}
		if p1 == nil {
			cur = p2
			p2  = p2.Next
		} else if p2 == nil {
			cur = p1
			p1  = p1.Next
		} else {
			if p1.Val < p2.Val {
				cur = p1
				p1  = p1.Next
			} else {
				cur = p2
				p2  = p2.Next
			}
		}

		if head == nil {
			head = cur
			tail = cur
		} else {
			tail.Next = cur
			tail = cur
		}
	}
	tail.Next = nil

	return head
}
