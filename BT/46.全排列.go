package BT

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	trace := make([]int, 0)
	backPermute(nums, trace, &ret)
	return ret
}

func backPermute(nums []int, trace []int, ret *[][]int) {
	if len(nums) == len(trace) {
		cpy := make([]int, len(trace))
		copy(cpy, trace)
		*ret = append(*ret, cpy)
		return
	}

	for i := 0; i < len(nums); i++ {
		isContain := false
		for _, val := range trace{
			if val == nums[i]{
				isContain = true
				break
			}
		}
		if isContain == true{
			continue
		}
		trace = append(trace, nums[i])
		backPermute(nums, trace, ret)
		trace = trace[:len(trace)-1]
	}
}
