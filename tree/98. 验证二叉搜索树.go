package tree

func isValidBST(root *TreeNode) bool {
	return recursionBST(root, nil, nil)
}

func recursionBST(node *TreeNode, minNode *TreeNode, maxNode *TreeNode) bool {
	if node == nil {
		return true
	}

	if minNode != nil && minNode.Val >= node.Val {
		return false
	}
	if maxNode != nil && maxNode.Val <= node.Val {
		return false
	}

	r1 := recursionBST(node.Left, minNode, node)
	r2 := recursionBST(node.Right, node, maxNode)

	return r1 && r2

}
