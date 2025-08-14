// https://codeforces.com/gym/103055/problem/A

/*
    From this problem I learned that the Go compiler can throw an error, not a warning,
    if a library is imported but not used.
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
	var blue, red, curr int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numsStr := strings.Fields(scanner.Text())
	for i := 0; i < 5; i++ {
		curr, _ = strconv.Atoi(numsStr[i])
		blue += curr
	}

	scanner.Scan()
	numsStr = strings.Fields(scanner.Text())
	for i := 0; i < 5; i++ {
		curr, _ = strconv.Atoi(numsStr[i])
		red += curr
	}

	if blue >= red {
		fmt.Println("Blue")
	} else {
		fmt.Println("Red")
	}
}