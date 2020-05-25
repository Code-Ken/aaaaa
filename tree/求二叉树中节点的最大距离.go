package tree

//case 1: 两个叶子节点都出现在左子树。find_max_len(root->pLeft);
//case 2: 两个叶子节点都出现在右子树。find_max_len(root->pRight);
//case 3: 一个出现在左子树，一个出现在右子树。distance(root->pLeft) + distance(root->pRight) + 2;其中，distance()计算子树中最远的叶子节点与根节点的距离，其实就是左子树的高度减1。

// 计算树的高度
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}


func findMaxLen(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := findMaxLen(root.Left)
	rightMax := findMaxLen(root.Right)
	lheight := 0
	rheight := 0

	if root.Left != nil{
		// 左子树最远的叶子节点与根节点的距离
		lheight = height(root.Left)
	}
	if root.Right != nil{
		// 右子树最远的叶子节点与根节点的距离
		rheight = height(root.Right)
	}

	if leftMax > rightMax && leftMax > lheight+rheight{
		return leftMax
	}

	if rightMax > leftMax && rightMax > lheight+rheight{
		return rightMax
	}
	return lheight+rheight

}
