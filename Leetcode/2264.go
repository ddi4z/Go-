// https://leetcode.com/problems/largest-3-same-digit-number-in-string/description/

func largestGoodInteger(num string) string {
	var res = ""
	for i := range len(num) - 2 {
		curr := num[i : i+3]
		if curr > res && num[i] == num[i+1] && num[i+1] == num[i+2] {
			res = curr
		}
	}
	return res
}

/*
   From this problem I learned:
   - Indexing a string in Go returns the byte value of the character
   - How to use string slicing in Go (it creates a substring referencing the same underlying data and stores the length)
   - How to convert a byte to a string using string(b)
*/
