package Other

type ListNode struct {
	Val  int
	Next *ListNode
}

//链表快排(分为三部分实现)
func sortList(head *ListNode) *ListNode {
	return quicksort(head)
}

func quicksort(head *ListNode) *ListNode {
	//出递归条件（没有节点或者只有一个节点时，不需要再分）
	if head == nil || head.Next == nil {
		return head
	}
	//遍历链表，将其分为三部分（三个哨兵节点）
	lhead := &ListNode{-1, nil} //小于基准值
	mhead := &ListNode{-1, nil} //等于基准值
	rhead := &ListNode{-1, nil} //大于基准值
	left, mid, right := lhead, mhead, rhead
	//基准值不一定非以最左边为基准
	pivot := head.Val //基准值（以最左边的节点为基准）
	//遍历链表
	for p := head; p != nil; p = p.Next {
		if p.Val < pivot { //小于给左边
			left.Next = p
			left = left.Next
		} else if p.Val > pivot { //大于给右边
			right.Next = p
			right = right.Next
		} else { //等于去中间
			mid.Next = p
			mid = mid.Next
		}
	}
	//将左中右节点值最后设为nil（防止后面还连着东西）
	left.Next = nil
	mid.Next = nil
	right.Next = nil
	//递归处理左右段
	lhead.Next = quicksort(lhead.Next)
	rhead.Next = quicksort(rhead.Next)
	//将左右段拼接起来
	//找到左端的尾节点
	getTail(lhead).Next = mhead.Next
	//找到中间的尾节点
	getTail(mhead).Next = rhead.Next
	return lhead.Next
}

//取得尾节点
func getTail(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}
