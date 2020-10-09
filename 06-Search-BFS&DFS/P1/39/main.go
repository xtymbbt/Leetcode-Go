package main

import "fmt"

func main() {
	fmt.Printf("result is : %#v", combinationSum([]int{2,3,5}, 8))
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
		mySort(&line)
		hasSame := false
		for _, re := range res {
			lineSame := true
			for j := range line {
				if line[j] != re[j]{
					lineSame = false
					break
				}
			}
			if lineSame {
				hasSame = true
			}
		}
		if hasSame {
			return
		}
		res = append(res, line)
		return
	}

	if nowCal > target {
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

func mySort(line *[]int) {
	if len(*line) == 0 {
		return
	}
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(*line); i++ {
		if (*line)[i] < (*line)[0] {
			left = append(left, (*line)[i])
		} else {
			right = append(right, (*line)[i])
		}
	}
	mySort(&left)
	mySort(&right)
	left = append(left, (*line)[0])
	*line = append(left, right...)
}