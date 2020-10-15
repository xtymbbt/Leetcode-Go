package main

import "fmt"

func main() {
	var x = []int{16,15,12,18,19,20,8,15,34,27,19,9,13,34,26,18,23,21,25,34,32,20,11,28,6,24,8,9,20,5,29,32,15,28,7,30,13,10,32,5,24,16,9,32,15,22,15,33,8,27,24,27,13,31,13,6,20,8}
	fmt.Printf("result is %#v\n", combinationSum2(x, 25))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort(&candidates)
	var res = make([][]int, 0)
	var line = make([]int, 0)
	bs(candidates, target, &res, line)
	return res
}

func bs(candidates []int, target int, res *[][]int, line []int) {
	if target == 0 && !contains(*res, line){
		*res = append(*res, line)
		return
	}
	for i, candidate := range candidates {
		var candidatesCP = make([]int, len(candidates))
		copy(candidatesCP, candidates)
		var lineCP = make([]int, len(line))
		copy(lineCP, line)
		lineCP = append(lineCP, candidate)
		nextTarget := target - candidate
		if nextTarget < 0 {
			continue
		}
		nextCandidates := candidatesCP[i+1:]
		if nextTarget >= 0 {
			bs(nextCandidates, nextTarget, res, lineCP)
		}
	}
	return
}

func contains(res [][]int, line []int) bool {
	//sort(&line)
	for _, re := range res {
		same := true
		//sort(&re)
		for i := range line {
			if line[i] != re[i] {
				same = false
				break
			}
		}
		if same {
			return true
		}
	}
	return false
}

func sort(array *[]int) {
	if len(*array) == 0 {
		return
	}
	base := (*array)[0]
	var (
		left = make([]int, 0)
		right = make([]int, 0)
	)

	for i:=1;i<len(*array);i++ {
		if (*array)[i] > base {
			left = append(left, (*array)[i])
		} else {
			right = append(right, (*array)[i])
		}
	}
	sort(&left)
	sort(&right)
	left = append(left, base)
	*array = append(left, right...)
	return
}