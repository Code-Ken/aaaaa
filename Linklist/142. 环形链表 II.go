package main

func detectCycle(head *ListNode) *ListNode {
	fastP, slowP := head, head

	for fastP != nil && fastP.Next != nil {
		fastP = fastP.Next.Next
		slowP = slowP.Next
		if fastP == slowP {
			fastP = head
			for {
				if fastP == slowP {
					return fastP
				}
				fastP = fastP.Next
				slowP = slowP.Next
			}
		}
	}
	return nil
}
