package main

func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p!=nil {
		val := p.Val
		for p.Next!=nil && p.Next.Val ==  val{
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head                             // 终止条件
	}
	head.Next  = deleteDuplicates(head.Next)     // 递归调用
	if head.Val == head.Next.Val {
		head = head.Next                      // 去重
	}
	return head
}

