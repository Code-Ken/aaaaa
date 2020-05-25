package Other

import "sort"

//按照官方来说就是先排序,再合并


func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{}
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged)-1]
		c := intervals[i]

		if c[0] > m[1] {
			merged = append(merged, c)
			continue
		}

		if c[1] > m[1] {
			m[1] = c[1]
		}
	}

	return merged
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	/* 选择排序 */
	for i := 0; i < len(intervals)-1; i++ {
		min := i
		for j := i + 1; j < len(intervals); j++ {
			if intervals[min][0] > intervals[j][0] {
				min = j
			}
		}
		intervals[i], intervals[min] = intervals[min], intervals[i]
	}
	var ret [][]int
	l, r := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if r >= intervals[i][0] {
			if r < intervals[i][1] {
				r = intervals[i][1]
			}
		} else {
			ret = append(ret, []int{l, r})
			l, r = intervals[i][0], intervals[i][1]
		}
	}
	ret = append(ret, []int{l, r})
	return ret
}

