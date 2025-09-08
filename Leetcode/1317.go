// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/

/*
	In Go you must always return something, you can use nil to avoid this.
*/

func check(num int) bool {
	if num == 0 {
		return false
	}
	for num > 0 {
		if num%10 == 0 {
			return false
		}
		num /= 10
	}
	return true
}

func getNoZeroIntegers(n int) []int {
	for a := range n {
		b := n - a
		if check(a) && check(b) {
			return []int{a, b}
		}
	}
	return nil
}
