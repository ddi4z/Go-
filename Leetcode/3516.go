// https://leetcode.com/problems/find-closest-person/

func findClosest(x int, y int, z int) int {
	var dx, dy int = x - z, y - z
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	if dx == dy {
		return 0
	}
	if dx > dy {
		return 2
	}
	return 1
}