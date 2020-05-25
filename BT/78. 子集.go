package BT

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	track := make([]int, 0)
	backtrackSubset(nums, 0, track, &ans)
	return ans
}

func backtrackSubset(nums []int, start int, track []int, ans *[][]int) {
	cpyTrack := make([]int, len(track))
	copy(cpyTrack, track)
	*ans = append(*ans, cpyTrack)


	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrackSubset(nums, i+1, track, ans)
		track = track[:len(track)-1]
	}
}

