package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	fmt.Println(g)
	fmt.Println(s)
	ans := 0
	for i, j := 0, 0;i < len(g) && j < len(s);i++ {
		if g[i] <= s[j] {
			ans++
		} else {
			i--
		}
		j++
	}
	return ans
}

func main() {
	fmt.Println(findContentChildren([]int{10, 9, 8, 7}, []int{5,6,7,8}))
}