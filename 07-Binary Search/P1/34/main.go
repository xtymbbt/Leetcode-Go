package _4

func searchRange(nums []int, target int) []int {
	return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := l + (r - l) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}

	if l == len(nums) || nums[l] != target {
		return -1
	}
	return l
}

func findRight(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := l + (r - l) / 2
		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	// l points to the first element this is greater than target.
	l--
	if l < 0 || nums[l] != target {
		return -1
	}
	return l
}
