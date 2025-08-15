// https://codeforces.com/problemset/problem/1783/A

/*
   Go does not have a built-in Set type.
   A common way to implement one is by using a map where the keys represent the elements
   (often with struct{} as the value to save memory).

   To print a slice without the surrounding brackets, you can use:
       fmt.Println(strings.Trim(fmt.Sprint(order), "[]"))
   In some cases, the spread operator can be used (order...)
   if the slice is of type []any or converted to it.

   The two main ways to sort a slice in reverse order are:
       - sort.Slice(distinct, func(i, j int) bool { return distinct[i] > distinct[j] })
       - sort.Sort(sort.Reverse(sort.IntSlice(distinct)))
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
	for _ = range t {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		freq := make(map[int]int)
		for i := range n {
			curr, _ := strconv.Atoi(line[i])
			freq[curr]++
		}
		distinct := []int{}
		for key := range freq {
			distinct = append(distinct, key)
		}
		if len(distinct) == 1 && n > 1 {
			fmt.Println("NO")
			continue
		} else {
			fmt.Println("YES")
		}
		sort.Slice(distinct, func(i, j int) bool {
			return distinct[i] > distinct[j]
		})
		order := []int{}
		for len(order) != n {
			for _, num := range distinct {
				if freq[num] > 0 {
					freq[num]--
					order = append(order, num)
				}
			}
		}

		fmt.Println(strings.Trim(fmt.Sprint(order), "[]"))
	}

}
