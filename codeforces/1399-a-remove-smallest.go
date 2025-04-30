// I AM A MUSLIM
// URL: https://codeforces.com/contest/1399/problem/A

package main

import (
	"fmt"
	"sort"
)

func main() {
	var test int
	fmt.Scan(&test)

	for t := 0; t < test; t++ {
		var n int
		fmt.Scan(&n)
		var a []int
		for i := 0; i < n; i++ {
			var v int
			fmt.Scan(&v)
			a = append(a, v)
		}
		sort.Ints(a)
		ans := 1
		for i := 0; i < n-1; i++ {
			if a[i+1]-a[i] > 1 {
				ans = 0
				break
			}
		}
		
		if ans == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}