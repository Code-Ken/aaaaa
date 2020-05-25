package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	var head *ListNode
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}
	if p1.Val < p2.Val {
		head = p1
		p1 = p1.Next
	} else {
		head = p2
		p2 = p2.Next
	}
	cur := head

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}
		cur = cur.Next
	}
	if p1 != nil {
		cur.Next = p1
	}
	if p2 != nil {
		cur.Next = p2
	}
	return head
}
