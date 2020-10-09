package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var (
		results [][]int
		route   []int
		size    int
	)
	size = len(candidates)
	sort.Ints(candidates)
	backtrack(&route, 0, 0, target, candidates, size, &results)
	return results
}

func backtrack(route *[]int, start int, totalSum int, target int, candidates []int, size int, results *[][]int) {

	if totalSum == target {
		res :=make([]int, len(*route))
		copy(res, *route)
		*results = append(*results, res)
		return
	}
	for i := start; i < size; i++ {
		totalSum += candidates[i]
		if totalSum > target {
			totalSum -= candidates[i]
			break
		}
		*route = append(*route, candidates[i])
		backtrack(route, i, totalSum, target, candidates, size, results)
		totalSum -= candidates[i]
		*route = (*route)[:len(*route)-1]

	}
	return
}
