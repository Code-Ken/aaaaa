package main

func rotateRight(head *ListNode, k int) *ListNode {
	//判断是否为空链表
	if head == nil {
		return head
	}

	//计算链表长度
	temp := head
	res := head
	l := 1
	for temp.Next != nil {
		l++
		temp = temp.Next
	}
	//k对长度取余，得到实际需要选择的次数，如果不需要旋转直接返回head
	k = k % l
	if k == 0 {
		return head
	}

	//找到需要打断的位置，把temp后面的直接给res
	temp = head
	for i:=0; i<l-k-1; i++ {
		temp = temp.Next
	}
	res = temp.Next
	temp.Next = nil

	//找到res目前的结尾，上面已经将head的结尾打断，把head接上res末尾就是正确答案
	temp = res
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = head

	return res
}
