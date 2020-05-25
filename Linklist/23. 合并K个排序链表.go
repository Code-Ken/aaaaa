package main


func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	for l > 1 {
		for i := 0; i < l/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[l-i-1])
		}
		l = (l + 1) / 2
	}
	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	res := tmp
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil {
		tmp.Next = l2
	} else {
		tmp.Next = l1
	}
	return res.Next
}
