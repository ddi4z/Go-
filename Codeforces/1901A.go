// https://codeforces.com/problemset/problem/1901/A

/*
	From this problem, I learned that Go will throw an error if a variable is declared but not used.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tests, _ := strconv.Atoi(scanner.Text())

	for t := 0; t < tests; t++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(parts[0])
		x, _ := strconv.Atoi(parts[1])

		scanner.Scan()
		numsStr := strings.Fields(scanner.Text())
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			nums[j], _ = strconv.Atoi(numsStr[j])
		}
		res := nums[0]
		for i := 1; i < n; i++ {
			res = max(res, nums[i]-nums[i-1])
		}
		res = max(res, 2*(x-nums[n-1]))
		fmt.Println(res)
	}
}
