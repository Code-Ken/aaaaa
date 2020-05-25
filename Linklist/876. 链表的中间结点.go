package main

func middleNode(head *ListNode) *ListNode {
	fastP,slowP := head,head
	for nil != fastP && nil != fastP.Next{
		fastP,slowP = fastP.Next.Next,slowP.Next
	}
	return slowP
}

