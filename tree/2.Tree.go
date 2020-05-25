package tree
//https://github.com/temporaries/leetcode/tree/master/templates/tree
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//2.1 中序遍历
func inorderTraversal(root *TreeNode) []int{
	res := make([]int,0)
	inorderHelper(root,&res)
	return res
}

//递归中序遍历
func inorderHelper(node *TreeNode,res *[]int){
	if node == nil{
		return
	}

	if node.Left != nil{
		inorderHelper(node.Left,res)
	}
	*res = append(*res,node.Val)
	if node.Right != nil{
		inorderHelper(node.Right,res)
	}
}

//非递归中序遍历
func inorderHelper2(root *TreeNode)[]int{
	res := make([]int,0)
	var st stack.Stack
	cur := root
	for cur != nil || st.Len() != 0{
		for cur != nil{
			st.Push(cur)
			cur = cur.Left
		}
		cur = st.Pop().(*TreeNode)
		res = append(res,cur.Val)
		cur = cur.Right
	}
	return res
}
//非递归中序遍历 数组栈实现
func inorderStack(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for 0 < len(stack) || root != nil { //root != nil 只为了第一次root判断，必须放最后
		for root != nil {
			stack = append(stack, root) //入栈
			root = root.Left            //移至最左
		}
		index := len(stack) - 1             //栈顶
		res = append(res, stack[index].Val) //中序遍历
		root = stack[index].Right           //右节点会进入下次循环，如果 =nil，继续出栈
		stack = stack[:index]               //出栈
	}
	return res
}

//2.1 前序遍历
func preorderTraversal(root *TreeNode) []int {
	res := make([]int,0)
	preorderHelper(root,&res)
	return res
}

//递归前序遍历
func preorderHelper(node *TreeNode,res *[]int){
	if node == nil{
		return
	}

	*res = append(*res,node.Val)
	preorderHelper(node.Left,res)
	preorderHelper(node.Right,res)
}

//非递归前序遍历
func preorderHelper2(node *TreeNode) []int{
	var st stack.Stack
	res := make([]int,0)

	cur := root

	for cur != nil ||st.Len()!=0{
		for cur != nil{
			res = append(res,node.Val)
			st.Push(cur)
			cur = cur.Left
		}
		cur = cur.Pop().(*TreeNode)
		cur = cur.Right
	}
	return res
}

//2.1 后序遍历
func postorderTraversal(root *TreeNode) []int {
	res := make([]int,0)

	postorderHelper(root,&res)

	return res
}

//递归后序遍历
func postorderHelper(node *TreeNode,res *[]int){
	if node == nil{
		return
	}
	if node.Left != nil{
		postorderHelper(node.Left,res)
	}
	if node.Right != nil{
		postorderHelper(node.Right,res)
	}
	*res = append(*res,node.Val)
}

//非递归后序遍历
func postorderHelper2(root *TreeNode) []int{
	res := make([]int,0)
	var st stack.Stack
	cur := root
	lastVisit := root
	for cur != nil || st.Len()!=0{
		for cur != nil{
			st.Push(cur)
			cur = cur.Left
		}
		cur = st.Peek().(*TreeNode)
		if cur.Right == nil || cur.Right == lastVisit{
			res = append(res,cur.Val)
			lastVisit = st.Pop().(*TreeNode)
			cur = nil
		}else{
			cur = cur.Right
		}
	}
	return res
}

// 翻转二叉树

func invertTree(root *TreeNode) *TreeNode{
	recurTree(root)
	return root
}

//递归翻转二叉树
func recurTree(node *TreeNode){
	if node == nil{
		return
	}
	node.Left,node.Right = node.Right,node.Left
	recurTree(node.Left)
	recurTree(node.Right)
}

//非递归翻转二叉树
func invertTree2(root *TreeNode) *TreeNode{
	var qu Queue
	if root == nil{
		return root
	}
	qu.Enqueue(root)

	for qu.Len() != 0{
		l := qu.Len()
		for i:=0;i<l;i++{
			node := qu.Dequeue().(*TreeNode)
			node.Left,node.Right = node.Right,node.Left
			if node.Left != nil{
				qu.Enqueue(node.Left)
			}
			if node.Right != nil{
				qu.Enqueue(node.Right)
			}
		}
	}
	return root
}

//N叉树的层序遍历
//递归 层序遍历
var res [][]int

func levelOrder(root *Node) [][]int{
	levelOrderHelper(root,0)
	return res
}

func levelOrderHelper(node *Node,level int){
	if node == nil{
		return
	}

	if len(res) == level{
		res = append(res,[]int{})
	}
	res[level] = append(res[level],root.Val)

	for _,cNode := range node.Children{
		levelOrderHelper(cNode,level+1)
	}

}

//非递归层序遍历
func levelOrder2(root *Node) [][]int{
	var qu Queue
	var ret [][]int

	if root == nil{
		return ret
	}

	qu.Enqueue(root)
	for qu.Len()!=0{
		n := qu.Len()
		tmp := make([]int,0)
		for i:=0;i<n;i++{
			curNode := qu.Dequeue().(*Node)
			tmp = append(tmp,curNode.Val)
			for _,childNode := range curNode.Children{
				qu.Enqueue(childNode)
			}
		}
		ret = append(res,tmp)
	}
	return ret
}
//非递归层序遍历2
func levelOrder3(root *TreeNode) [][]int {
	res := make([][]int,0)
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})
		for counter > 0 {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}
