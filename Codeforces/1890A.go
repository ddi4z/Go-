// https://codeforces.com/problemset/problem/1890/A

/*
	I learned how to use maps in Go.
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

		scanner.Scan()
		line := strings.Fields(scanner.Text())

		m := make(map[int]int)
		for i := range n {
			curr, _ := strconv.Atoi(line[i])
			m[curr]++
		}

		if len(m) == 1 {
			fmt.Println("YES")
		} else if len(m) >= 3 {
			fmt.Println("NO")
		} else {
			keys := []int{}
			for k := range m {
				keys = append(keys, k)
			}
			if (m[keys[0]]-m[keys[1]]) <= 1 && (m[keys[0]]-m[keys[1]]) >= -1 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}

		}

	}
}
