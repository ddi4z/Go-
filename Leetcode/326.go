//https://leetcode.com/problems/power-of-three/description/

/*
	From this problem I learned that dividing two integers in Go produces an integer result (truncating any decimal).
*/

func isPowerOfThree(n int) bool {
	for n > 1 {
		if (n % 3) != 0 {
			return false
		}
		n /= 3
	}
	return n == 1
}