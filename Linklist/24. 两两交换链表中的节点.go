package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dumpHead := &ListNode{Val: 0, Next: head}
	p := dumpHead
	cur := dumpHead.Next

	for cur != nil && cur.Next != nil {
		//交换
		p.Next = cur.Next
		cur.Next = cur.Next.Next
		p.Next.Next = cur

		//后移指针
		p = cur
		cur = cur.Next
	}

	return dumpHead.Next
}
