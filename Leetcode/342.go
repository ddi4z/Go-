// https://leetcode.com/problems/power-of-four/description/

func isPowerOfFour(n int) bool {
	for n > 1 {
		if (n % 4) != 0 {
			return false
		}
		n /= 4
	}
	return n == 1
}