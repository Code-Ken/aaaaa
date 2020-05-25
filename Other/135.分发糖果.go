package Other

/*
一开始大家都是1，所以最少是len(n)，这个在后面的结果加上去就行了。
建一个切片存放每个结果，res[i]表示第i个孩子能分多少个。

从左向右遍历，如果右边大于左边，则右边=左边+1；

从右倒数第二个向左遍历，如果左边的大于右边且分数小于左边小于等于右边，则左边=右边+1

最后，累加res+len(n)

*/

func candy(ratings []int) int {
	l := len(ratings)
	res := make([]int, len(ratings))

	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}
	}

	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && res[i] <= res[i+1] {
			res[i] = res[i+1] + 1
		}
	}

	sum := l
	for _, v := range res {
		sum += v
	}

	return sum
}
