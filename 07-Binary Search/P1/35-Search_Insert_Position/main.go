package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (right - left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// find left
			right = mid
		} else { // if nums[mid] < target
			// find right
			left = mid + 1
		}
	}
	return left
}

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
}