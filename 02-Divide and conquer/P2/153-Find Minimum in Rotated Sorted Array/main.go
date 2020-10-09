package notMain

func findMin(nums []int) int {
	var min = nums[0]
	for _, num := range nums {if min > num {return num}}
	return min
}