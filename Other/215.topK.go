package main

func findKthLargest(nums []int, k int) int {
	nums = heapSort(nums)
	return nums[k-1]
}



func heapSort(nums []int) []int {
	lens := len(nums)
	// 先建立一个堆
	buildHeap(nums,lens)
	// 上面构建堆后,其实已经满足了基本的 最大堆或者最小堆
	// 此时根节点肯定是最大值或者最小值了
	// 从数组最后挨个往前遍历
	for i:=lens-1;i>=0;i-- {
		// 构成堆后的数组的第一个值和最后一个值进行替换
		// 然后把 数组 i 个值之后的数之外的数 继续构建堆
		// 循环往复
		swap(nums,0,i)
		lens -= 1
		heap(nums,0,lens)
	}
	return nums
}

// 构建一个堆
func buildHeap(nums []int,lens int) {
	// 一般数组长度的一半 很可能就不是叶子节点了,但可以确定的是 数组长度一半后面的肯定都是叶子节点
	// 叶子节点是最底下的节点了,所以不需要往下置换了,所以从 长度/2开始算
	for i := lens/2;i>=0;i-- {
		heap(nums,i,lens)
	}
}

func heap(nums []int,i,lens int) {
	// 每个父节点的左节点 是当前数组下标i的 2*i + 1
	// 每个父节点的右节点 是当前数组下标i的 2*i + 2
	left := 2*i + 1
	right := 2*i + 2

	// 咱们做最小堆,也就是小的数组排在后面
	// 假设当前的 父节点是最小
	min := i
	// 当 父节点的左子节点 小于 数组总长度 且 左子节点的值小于父节点的时候,所以这个小堆里 左子节点是最小的(暂时不替换)
	// 为什么要 父节点的左子节点 小于 数组总长度, 因为如果父节点的左子节点比 数组总长度还长,那说明 超出了数组范围了..说明该父节点其实没有左子节点
	if left < lens && nums[left] < nums[min] {
		min = left
	}
	// 右子节点跟上面左子节点一个道理
	if right < lens && nums[right] < nums[min] {
		min = right
	}

	// 如果通过上面与左右子节点相比,最小值已经不是当初的父节点了
	// 那么把最小的节点与父节点进行变换,变换后可能会影响该父节点的子节点的子节点,所以还得往下继续对比比换一下
	if min != i {
		swap(nums,min,i)
		heap(nums,min,lens)
	}

}


func swap(arr []int,m,n int) []int {
	arr[m],arr[n] = arr[n],arr[m]
	return arr
}

