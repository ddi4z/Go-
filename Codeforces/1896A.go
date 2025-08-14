// https://www.onlinegdb.com/online_go_compiler

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
	tests, _ := strconv.Atoi(scanner.Text())

	for t := 0; t < tests; t++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(parts[0])

		scanner.Scan()
		line := strings.Fields(scanner.Text())
		nums := make([]int, n)

		for i := range n {
			nums[i], _ = strconv.Atoi(line[i])
		}
		for _ = range n {
			for i := 1; i < n-1; i++ {
				if nums[i] > max(nums[i-1], nums[i+1]) {
					nums[i], nums[i+1] = nums[i+1], nums[i]
				}
			}
		}

		if sort.IntsAreSorted(nums) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
