package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	zero := &ListNode{
		Next: head,
	}
	left := zero
	right := zero
	for i := 0; i <= n; i++ {
		right = right.Next
	}
	for right != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return zero.Next
}
