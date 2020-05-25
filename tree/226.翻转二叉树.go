package tree

func invertTree(root *TreeNode) *TreeNode {
	invertTreeRecursion(root)
	return root
}

func invertTreeRecursion(node *TreeNode) {
	if node == nil {
		return
	}

	node.Right, node.Left = node.Left, node.Right

	invertTreeRecursion(node.Left)
	invertTreeRecursion(node.Right)
}

