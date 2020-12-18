package main

func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	maxIndex := findMaxIndex(height)
	leftArea := calcLeft(height[:maxIndex])
	rightArea := calcRight(height[maxIndex+1:])
	return leftArea+rightArea
}

func calcLeft(nums []int) int {
	if len(nums) < 2{
		return 0
	}
	maxIndex := findMaxIndex(nums)
	rightArea := 0
	for i := maxIndex+1; i < len(nums); i++ {
		rightArea += nums[maxIndex] - nums[i]
	}
	return rightArea + calcLeft(nums[:maxIndex])
}

func calcRight(nums []int) int {
	if len(nums) < 2{
		return 0
	}
	maxIndex := findMaxIndex(nums)
	leftArea := 0
	for i := 0; i < maxIndex; i++ {
		leftArea += nums[maxIndex] - nums[i]
	}
	return leftArea + calcRight(nums[maxIndex+1:])
}

func findMaxIndex(nums []int) int {
	max := 0
	index := 0
	for i, num := range nums {
		if num > max {
			max = num
			index = i
		}
	}
	return index
}