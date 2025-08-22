// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/description/

func minimumArea(grid [][]int) int {
	minCol := len(grid[0]) + 1
	maxCol := -1
	minRow := len(grid) + 1
	maxRow := -1
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 1 {
				minCol = min(minCol, j)
				maxCol = max(maxCol, j)
				minRow = min(minRow, i)
				maxRow = max(maxRow, i)
			}
		}
	}
	return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}