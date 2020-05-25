package Other

import "math"

func findKthNumber(n int, k int) int {

	currentNode := 1
	prefix := 1

	for currentNode < k {
		count := getCount(prefix, n)
		if (currentNode + count) > k {
			prefix *= 10
			currentNode++
		} else {
			prefix++
			currentNode += count
		}
	}
	return prefix
}

func getCount(prefix int, n int) int {
	count := 0
	cur := prefix
	next := prefix + 1

	for cur <= n {
		count += int(math.Min(float64(next), float64(n+1))) - cur
		cur *= 10
		next *= 10
	}

	return count
}
