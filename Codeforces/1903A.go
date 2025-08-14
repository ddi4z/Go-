//https://codeforces.com/problemset/problem/1903/A

/*
   From this problem I learned:
   - To declare dynamically sized arrays, you need to use nums := make([]int, n)
   - I learned how to safely read input lines
   - I learned how to use the sort package to check if an array is sorted
   - I learned how to use the strconv package to convert strings to integers
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(parts[0])
		k, _ := strconv.Atoi(parts[1])

		scanner.Scan()
		numsStr := strings.Fields(scanner.Text())
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			nums[j], _ = strconv.Atoi(numsStr[j])
		}

		if k == 1 && !sort.IntsAreSorted(nums) {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
