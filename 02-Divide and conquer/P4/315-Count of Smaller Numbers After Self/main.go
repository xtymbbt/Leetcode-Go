package main

func main() {
	
}

var res []int

func countSmaller(nums []int) []int {
	res = make([]int, len(nums))

	for i, _ := range nums {
		res[i] = calcNum(nums, i)
	}

	return res
}

func calcNum(nums []int, position int) int {
	sum := 0
	for i := position+1; i < len(nums); i++ {
		if nums[i] < nums[position] {
			sum++
		}
	}
	return sum
}