// https://leetcode.com/problems/count-square-submatrices-with-all-ones/description/

/*
	In go, a number inside an if statement is not treated as a boolean.
*/

func countSquares(matrix [][]int) int {
	var result int
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			if matrix[i][j] == 1 {
				if j == 0 || i == 0 {
					result += 1
					continue
				}
				side := 1 + min(matrix[i-1][j-1], matrix[i-1][j], matrix[i][j-1])
				matrix[i][j] = side
				result += side
			}
		}
	}
	return result
}