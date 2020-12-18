package main

func maxArea(height []int) int {
	l := 0
	r := len(height)-1
	max := 0
	for l != r {
		diff := height[l] - height[r]
		if diff > 0{
			area := height[r] * (r - l)
			if area > max {
				max = area
			}
			r--
		} else {
			area := height[l] * (r - l)
			if area > max {
				max = area
			}
			l++
		}
	}
	return max
}
