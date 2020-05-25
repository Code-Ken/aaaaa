package BT


func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	trace := make([]int, 0)
	backCombine(n, k, 1, trace, &ret)
	return ret

}

func backCombine(n int, k int, start int, trace []int, ret *[][]int) {
	if len(trace) == k {
		cpy := make([]int,k)
		copy(cpy,trace)
		*ret = append(*ret, cpy)
	}

	for i := start; i <= n; i++ {
		trace = append(trace, i)
		backCombine(n, k, i+1, trace, ret)
		trace = trace[:len(trace)-1]
	}
}