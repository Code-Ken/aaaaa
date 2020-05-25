package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := l1
	var carry int
	for {
		l1.Val += l2.Val + carry
		carry = l1.Val / 10
		l1.Val = l1.Val % 10
		if l2.Next == nil {
			if carry == 0 {
				break
			}
			l2.Next = &ListNode{}
		}
		if l1.Next == nil {
			if carry == 0 {
				l1.Next = l2.Next
				break
			}
			l1.Next = &ListNode{}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l
}
