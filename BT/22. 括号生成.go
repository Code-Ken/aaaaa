package BT

func generateParenthesis(n int) []string {
	ret := make([]string, 0, 0)
	helper(0, 0, n, "", &ret)
	return ret
}

func helper(left int, right int, total int, s string, ret *[]string) {
	if left == right && left == total {
		*ret = append(*ret, s)
		return
	}

	if left < total {
		helper(left+1, right, total, s+"(", ret)
	}

	if right < left {
		helper(left, right+1, total, s+")", ret)
	}
}

//=========================================

func generateParenthesis(n int) []string {

	var result = make([]string, 0, 0)

	if n == 0 {
		return result
	}

	parenthesisRecursion(0, "", n<<1, 0, 0, &result)
	return result
}

func parenthesisRecursion(currentLen int, curStr string, l int, leftC int, rightC int, result *[]string) {
	if len(curStr) == l {
		*result = append(*result, curStr)
		return
	}

	//process current ...

	//drill down
	if leftC < l>>1 {
		parenthesisRecursion(currentLen+1, curStr+"(", l, leftC+1, rightC, result)
	}
	if rightC < leftC {
		parenthesisRecursion(currentLen+1, curStr+")", l, leftC, rightC+1, result)
	}
}
