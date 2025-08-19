// https://leetcode.com/problems/number-of-zero-filled-subarrays/

/*
	You can convert int to int64 using int64()
	Int is 32 bits and int64 is 64 bits
*/

func zeroFilledSubarray(nums []int) int64 {
	var result int64 = 0
	i := 0
	for i < len(nums) {
		if nums[i] == 0 {
			j := i
			for j < len(nums) && nums[j] == 0 {
				result += int64(j - i + 1)
				j++
			}
			i = j
		} else {
			i++
		}
	}
	return result
}
