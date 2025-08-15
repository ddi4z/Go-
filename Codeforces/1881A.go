// https://codeforces.com/problemset/problem/1881/A

/*
	I learned how to use classic while loops in Go.
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
	t, _ := strconv.Atoi(scanner.Text())
	for _ = range t {
		scanner.Scan()
		_ = strings.Fields(scanner.Text())
		scanner.Scan()
		x := scanner.Text()
		scanner.Scan()
		s := scanner.Text()
		var res int
		for len(x) < len(s) {
			x += x
			res++
		}
		var found bool
		for i := range len(x) - len(s) + 1 {
			if x[i:i+len(s)] == s {
				found = true
			}
		}
		if found {
			fmt.Println(res)
			continue
		}
		x += x
		res++
		for i := range len(x) - len(s) + 1 {
			if x[i:i+len(s)] == s {
				found = true
			}
		}
		if found {
			fmt.Println(res)
			continue
		}
		fmt.Println(-1)
	}

}
