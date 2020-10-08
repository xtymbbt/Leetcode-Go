package main

func main() {
	combinationSum([]int{2,3,6,7}, 7)
}

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	line := make([]int, 0)
	dfs(candidates, target, line, 0)
	return res
}

func dfs(candidates []int, target int, line []int, nowCal int) {

	if nowCal == target {
		res = append(res, line)
		return
	}

	for _, candidate := range candidates {
		var lineCP = make([]int, len(line))
		copy(lineCP, line)
		lineCP = append(lineCP, candidate)
		dfs(candidates, target, lineCP, nowCal+candidate)
	}
	return
}
