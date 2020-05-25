package DP

/*

Map<TreeNode, Integer> memo = new HashMap<>();
public int rob(TreeNode root) {
    if (root == null) return 0;
    // 利用备忘录消除重叠子问题
    if (memo.containsKey(root))
        return memo.get(root);
    // 抢，然后去下下家
    int do_it = root.val
        + (root.left == null ?
            0 : rob(root.left.left) + rob(root.left.right))
        + (root.right == null ?
            0 : rob(root.right.left) + rob(root.right.right));
    // 不抢，然后去下家
    int not_do = rob(root.left) + rob(root.right);

    int res = Math.max(do_it, not_do);
    memo.put(root, res);
    return res;
}
*/


func rob(root *TreeNode) int {
	ret := subInterRob(root)
	return maxRob3(ret[0], ret[1])
}

func subInterRob(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	result := []int{0, 0}
	left := subInterRob(root.Left)
	right := subInterRob(root.Right)

	/* 返回一个大小为 2 的数组 arr
	arr[0] 表示不抢 root 的话，得到的最大钱数
	arr[1] 表示抢 root 的话，得到的最大钱数
	*/
	result[0] = maxRob3(left[0], left[1]) + maxRob3(right[0], right[1])
	result[1] = root.Val + left[0] + right[0]
	return result
}

func maxRob3(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
