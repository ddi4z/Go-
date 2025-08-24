// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

func longestSubarray(nums []int) int {
	var groups = make([]int, 0)
	var curr, res int

	for _, num := range nums {
		if num == 1 {
			curr += 1
		} else {
			groups = append(groups, curr)
			curr = 0
			groups = append(groups, 0)
		}
	}
	if curr > 0 {
		groups = append(groups, curr)
	}
	if len(groups) > 0 {
		res = groups[0]
		if groups[0] == len(nums) {
			res = groups[0] - 1
		}
	}
	for i := range (len(groups)) - 2 {
		res = max(res, groups[i]+groups[i+2])
	}
	return res
}