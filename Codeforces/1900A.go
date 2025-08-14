// https://codeforces.com/problemset/problem/1900/A

/*
	I learned how to append elements to a slice in Go. (The method returns a new slice)
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
		line := scanner.Text()

		obstacles := []int{-1}
		for i := 0; i < n; i++ {
			if line[i] == '#' {
				obstacles = append(obstacles, i)
			}
		}
		obstacles = append(obstacles, n)
		var found bool
		var res int
		for i := 1; (i < len(obstacles)) && !found; i++ {
			numElements := obstacles[i] - obstacles[i-1] - 1
			if (numElements) >= 3 {
				fmt.Println(2)
				found = true
			}
			res += numElements

		}
		if !found {
			fmt.Println(res)
		}

	}
}
